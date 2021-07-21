package queue

import (
	"context"
)

type (
	// Follower of the Leader.
	Follower interface {
		// Start following the Leader.
		Start(context.Context)
		// Stop following the Leader.
		Stop()
	}

	follower struct {
		leader   Leader
		jobQueue chan Job
		quit     chan struct{}
	}
)

// NewFollower creates a new Follower instance.
func NewFollower(leader Leader) Follower {
	return &follower{
		leader:   leader,
		jobQueue: make(chan Job),
		quit:     make(chan struct{}),
	}
}

// Start the Follower's duty.
func (f follower) Start(controlCtx context.Context) {
	go func(controlCtx context.Context) {
		for {
			f.leader.AddToFollowersPool(f.jobQueue)

			select {
			case job := <-f.jobQueue:

				dataTransfer := NewDataTransfer().
					Set(PayloadKey, job.Payload())

				controlCtx = context.WithValue(controlCtx, DataTransfer, dataTransfer)

				if err := job.Workflow().Execute(controlCtx); err != nil {
					f.leader.ReportError(err)
				}

				f.leader.ReportCompletion()
			case <-controlCtx.Done():
			case <-f.quit:
				return
			}
		}
	}(controlCtx)
}

// Stop the Follower's duty.
func (f follower) Stop() {
	go func() {
		f.quit <- struct{}{}
	}()
}
