package walletnode

import (
	"bufio"
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/pastelnetwork/gonode/common/errors"
	"github.com/pastelnetwork/gonode/common/log"
	"github.com/pastelnetwork/gonode/common/random"
	pb "github.com/pastelnetwork/gonode/proto/walletnode"
	"github.com/pastelnetwork/gonode/supernode/node/grpc/server/services/common"
	"github.com/pastelnetwork/gonode/supernode/services/artworkregister"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

// RegisterArtowrk represents grpc service for registration artowrk.
type RegisterArtowrk struct {
	pb.UnimplementedRegisterArtowrkServer

	*common.RegisterArtowrk
	workDir string
}

// Health implements supernode.RegisterArtowrkServer.Health()
func (service *RegisterArtowrk) Health(stream pb.RegisterArtowrk_HealthServer) error {
	ctx := stream.Context()

	task, err := service.TaskFromMD(ctx)
	if err != nil {
		return err
	}
	defer task.Cancel()

	peer, _ := peer.FromContext(ctx)
	log.WithContext(ctx).WithField("addr", peer.Addr).Debugf("Helath stream")
	defer log.WithContext(ctx).WithField("addr", peer.Addr).Debugf("Helath stream closed")

	go func() {
		defer task.Cancel()
		for {
			if _, err := stream.Recv(); err != nil {
				return
			}
		}
	}()

	<-task.Done()
	return nil
}

// Handshake implements supernode.RegisterArtowrkServer.Handshake()
func (service *RegisterArtowrk) Handshake(ctx context.Context, req *pb.HandshakeRequest) (*pb.HandshakeReply, error) {
	log.WithContext(ctx).WithField("req", req).Debugf("Handshake request")

	task := service.NewTask(ctx)
	if err := task.Handshake(ctx, req.IsPrimary); err != nil {
		return nil, err
	}

	resp := &pb.HandshakeReply{
		ConnID: task.ID,
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("Handshake response")
	return resp, nil
}

// AcceptedNodes implements supernode.RegisterArtowrkServer.AcceptedNodes()
func (service *RegisterArtowrk) AcceptedNodes(ctx context.Context, req *pb.AcceptedNodesRequest) (*pb.AcceptedNodesReply, error) {
	log.WithContext(ctx).WithField("req", req).Debugf("AcceptedNodes request")
	task, err := service.TaskFromMD(ctx)
	if err != nil {
		return nil, err
	}

	nodes, err := task.AcceptedNodes(ctx)
	if err != nil {
		return nil, err
	}

	var peers []*pb.AcceptedNodesReply_Peer
	for _, node := range nodes {
		peers = append(peers, &pb.AcceptedNodesReply_Peer{
			NodeID: node.Key,
		})
	}

	resp := &pb.AcceptedNodesReply{
		Peers: peers,
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("AcceptedNodes response")
	return resp, nil
}

// ConnectTo implements supernode.RegisterArtowrkServer.ConnectTo()
func (service *RegisterArtowrk) ConnectTo(ctx context.Context, req *pb.ConnectToRequest) (*pb.ConnectToReply, error) {
	log.WithContext(ctx).WithField("req", req).Debugf("ConnectTo request")
	task, err := service.TaskFromMD(ctx)
	if err != nil {
		return nil, err
	}

	if err := task.ConnectTo(ctx, req.NodeID, req.ConnID); err != nil {
		return nil, err
	}

	resp := &pb.ConnectToReply{}
	log.WithContext(ctx).WithField("resp", resp).Debugf("ConnectTo response")
	return resp, nil
}

// SendImage implements supernode.RegisterArtowrkServer.SendImage()
func (service *RegisterArtowrk) SendImage(stream pb.RegisterArtowrk_SendImageServer) error {
	ctx := stream.Context()

	task, err := service.TaskFromMD(ctx)
	if err != nil {
		return err
	}

	fileID, _ := random.String(16, random.Base62Chars)
	filename := filepath.Join(service.workDir, fileID)

	file, err := os.Create(filename)
	if err != nil {
		return errors.Errorf("failed to open file %q: %w", filename, err)
	}

	// TODO: remove file at the end of work
	// defer func() {
	// 	os.Remove(filename)
	// 	log.WithContext(ctx).Debugf("Removed temp file %a", filename)
	// }()

	defer file.Close()
	log.WithContext(ctx).Debugf("Created temp file %q for uploading image", filename)

	wr := bufio.NewWriter(file)

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			if status.Code(err) == codes.Canceled {
				return errors.New("connection closed")
			}
			return errors.Errorf("failed to receive SendImage: %w", err)
		}

		if _, err := wr.Write(req.Payload); err != nil {
			return errors.Errorf("failed to write to file %q: %w", filename, err)
		}
	}

	// TODO: pass filename to the task
	_ = task

	resp := &pb.SendImageReply{}
	if err := stream.SendAndClose(resp); err != nil {
		return errors.Errorf("failed to send SendImage response: %w", err)
	}
	log.WithContext(ctx).WithField("resp", resp).Debugf("SendImage response")
	return nil
}

// Desc returns a description of the service.
func (service *RegisterArtowrk) Desc() *grpc.ServiceDesc {
	return &pb.RegisterArtowrk_ServiceDesc
}

// NewRegisterArtowrk returns a new RegisterArtowrk instance.
func NewRegisterArtowrk(service *artworkregister.Service, workDir string) *RegisterArtowrk {
	return &RegisterArtowrk{
		RegisterArtowrk: common.NewRegisterArtowrk(service),
		workDir:         workDir,
	}
}
