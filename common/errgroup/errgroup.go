package errgroup

import (
	"context"

	"github.com/pastelnetwork/gonode/common/errors"
	"golang.org/x/sync/errgroup"
)

// A Group is a collection of goroutines working on subtasks that are part of the same overall task.
type Group struct {
	*errgroup.Group
}

// Go calls the given function in a new goroutine and tries to recover from panics.
func (g *Group) Go(fn func() error) {
	g.Group.Go(func() (err error) {
		defer errors.Recover(func(rec error) { err = rec })
		return fn()
	})
}

// WithContext returns a new Group and an associated Context derived from ctx.
func WithContext(ctx context.Context) (*Group, context.Context) {
	group, ctx := errgroup.WithContext(ctx)
	return &Group{group}, ctx
}
