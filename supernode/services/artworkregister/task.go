package artworkregister

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/common/random"
	"github.com/pastelnetwork/gonode/supernode/services/artworkregister/state"
	"golang.org/x/sync/errgroup"
)

// Task is the task of registering new artwork.
type Task struct {
	*Service

	ID        string
	State     *state.State
	ImagePath string

	acceptMu    sync.Mutex
	accpetNodes Nodes

	connectNode *Node

	actCh chan func(ctx context.Context) error

	doneMu sync.Mutex
	doneCh chan struct{}
}

// Run starts the task
func (task *Task) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer task.Cancel()

	group, ctx := errgroup.WithContext(ctx)
	for {
		select {
		case <-ctx.Done():
		case <-task.Done():

		case action := <-task.actCh:
			group.Go(func() (err error) {
				defer errors.Recover(func(recErr error) { err = recErr })
				return action(ctx)
			})
			continue
		}
		break
	}
	cancel()
	return group.Wait()
}

// Cancel stops the task, which causes all connections associated with that task to be closed.
func (task *Task) Cancel() {
	task.doneMu.Lock()
	defer task.doneMu.Unlock()

	select {
	case <-task.Done():
		return
	default:
		log.WithPrefix(fmt.Sprintf("%s-%s", logPrefix, task.ID)).Debug("Task canceled")
		close(task.doneCh)
	}
}

// Done returns a channel when the task is canceled.
func (task *Task) Done() <-chan struct{} {
	return task.doneCh
}

// Session is handshake wallet to supernode
func (task *Task) Session(ctx context.Context, isPrimary bool) error {
	ctx = task.context(ctx)

	if err := task.requiredLatestStatus(state.StatusTaskStarted); err != nil {
		return err
	}

	if isPrimary {
		log.WithContext(ctx).Debugf("Acts as primary node")
		task.State.Update(ctx, state.NewStatus(state.StatusPrimaryMode))
		return nil
	}

	log.WithContext(ctx).Debugf("Acts as secondary node")
	task.State.Update(ctx, state.NewStatus(state.StatusSecondaryMode))
	return nil
}

// AcceptedNodes waits for connection supernodes, as soon as there is the required amount returns them.
func (task *Task) AcceptedNodes(ctx context.Context) (Nodes, error) {
	ctx = task.context(ctx)

	if err := task.requiredLatestStatus(state.StatusPrimaryMode); err != nil {
		return nil, err
	}
	log.WithContext(ctx).Debugf("Waiting for supernodes to connect")

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-task.State.Updated():
		if err := task.requiredLatestStatus(state.StatusAcceptedNodes); err != nil {
			return nil, err
		}
		return task.accpetNodes, nil
	}
}

// SessionNode accepts secondary node
func (task *Task) SessionNode(ctx context.Context, nodeID string) error {
	ctx = task.context(ctx)

	task.acceptMu.Lock()
	defer task.acceptMu.Unlock()

	if err := task.requiredLatestStatus(state.StatusPrimaryMode); err != nil {
		return err
	}

	if node := task.accpetNodes.ByID(nodeID); node != nil {
		return errors.Errorf("node %q is already registered", nodeID)
	}

	node, err := task.pastelNodeByExtKey(ctx, nodeID)
	if err != nil {
		return err
	}
	task.accpetNodes.Add(node)

	log.WithContext(ctx).WithField("nodeID", nodeID).Debugf("Accept secondary node")

	if len(task.accpetNodes) >= task.config.NumberConnectedNodes {
		task.State.Update(ctx, state.NewStatus(state.StatusAcceptedNodes))
	}
	return nil
}

// ConnectTo connects to primary node
func (task *Task) ConnectTo(_ context.Context, nodeID, sessID string) error {
	if err := task.requiredLatestStatus(state.StatusSecondaryMode); err != nil {
		return err
	}

	task.actCh <- func(ctx context.Context) error {
		ctx = task.context(ctx)

		node, err := task.pastelNodeByExtKey(ctx, nodeID)
		if err != nil {
			return err
		}

		if err := node.connect(ctx); err != nil {
			return err
		}

		if err := node.Session(ctx, task.config.PastelID, sessID); err != nil {
			return err
		}

		task.connectNode = node
		task.State.Update(ctx, state.NewStatus(state.StatusConnectedToNode))
		return nil
	}
	return nil
}

// UploadImage uploads an image
func (task *Task) UploadImage(_ context.Context, filename string) error {
	if err := task.requiredLatestStatus(state.StatusAcceptedNodes, state.StatusConnectedToNode); err != nil {
		return err
	}

	task.actCh <- func(ctx context.Context) error {
		ctx = task.context(ctx)

		task.ImagePath = filename
		task.State.Update(ctx, state.NewStatus(state.StatusImageUploaded))

		<-ctx.Done()

		if err := os.Remove(filename); err != nil {
			return errors.Errorf("failed to remove temp file %q: %w", filename, err)
		}
		log.WithContext(ctx).Debugf("Removed temp file %q", filename)
		return nil
	}
	return nil
}

func (task *Task) pastelNodeByExtKey(ctx context.Context, nodeID string) (*Node, error) {
	masterNodes, err := task.pastelClient.MasterNodesTop(ctx)
	if err != nil {
		return nil, err
	}

	for _, masterNode := range masterNodes {
		if masterNode.ExtKey != nodeID {
			continue
		}
		node := &Node{
			client:  task.Service.nodeClient,
			ID:      masterNode.ExtKey,
			Address: masterNode.ExtAddress,
		}
		return node, nil
	}

	return nil, errors.Errorf("node %q not found", nodeID)
}

func (task *Task) requiredLatestStatus(statusTypes ...state.StatusType) error {
	latest := task.State.Latest()
	if latest == nil {
		return errors.New("not found latest status")
	}

	for _, statusType := range statusTypes {
		if latest.Type == statusType {
			return nil
		}
	}
	return errors.Errorf("wrong order, current task status %q, ", latest.Type)
}

func (task *Task) context(ctx context.Context) context.Context {
	return log.ContextWithPrefix(ctx, fmt.Sprintf("%s-%s", logPrefix, task.ID))
}

// NewTask returns a new Task instance.
func NewTask(service *Service) *Task {
	taskID, _ := random.String(8, random.Base62Chars)

	return &Task{
		Service: service,
		ID:      taskID,
		State:   state.New(state.NewStatus(state.StatusTaskStarted)),
		doneCh:  make(chan struct{}),
		actCh:   make(chan func(ctx context.Context) error),
	}
}
