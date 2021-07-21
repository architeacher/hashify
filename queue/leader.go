package queue

import (
	"context"
	"sync"
)

type (
	// Leader of one or more Follower.
	Leader interface {
		// Run as fast as he can.
		Run(context.Context)
		// Stop the running.
		Stop()
		// AddToFollowersPool an idle follower.
		AddToFollowersPool(chan Job)
		// ReportCompletion by a Follower.
		ReportCompletion()
		// ReportError to the error queue.
		ReportError(error)
		// Wait on the Leader to finish.
		Wait()
	}

	leader struct {
		jobsQueue         chan Job
		errorsQueue       chan error
		team              []Follower
		followersChanPool chan chan Job
		wg                sync.WaitGroup
		finished          chan struct{}
		concurrency       uint
	}
)

// NewLeader creates a new Leader instance.
func NewLeader(jobsQueue chan Job, errorsQueue chan error, concurrency uint) Leader {
	return &leader{
		jobsQueue:         jobsQueue,
		errorsQueue:       errorsQueue,
		team:              make([]Follower, concurrency),
		followersChanPool: make(chan chan Job, concurrency),
		finished:          make(chan struct{}),
		concurrency:       concurrency,
	}
}

func (l *leader) Run(controlCtx context.Context) {
	l.buildTeam(controlCtx)
	l.consume()
}

func (l *leader) AddToFollowersPool(idleFollowerChan chan Job) {
	l.followersChanPool <- idleFollowerChan
}

func (l *leader) ReportError(err error) {
	l.errorsQueue <- err
}

func (l *leader) buildTeam(controlCtx context.Context) {
	for n := uint(0); n < l.concurrency; n++ {
		l.team[n] = NewFollower(l)
		l.team[n].Start(controlCtx)
	}
}

func (l *leader) consume() {
	go func() {
		for job := range l.jobsQueue {
			l.wg.Add(1)

			go func(job Job) {
				jobChan := l.getIdleFollowerChan()
				jobChan <- job
			}(job)
		}
		l.wg.Wait()
		l.Stop()
		l.finished <- struct{}{}
	}()
}

func (l *leader) ReportCompletion() {
	l.wg.Done()
}

func (l *leader) Wait() {
	<-l.finished
}

func (l *leader) getIdleFollowerChan() chan Job {
	return <-l.followersChanPool
}

func (l *leader) Stop() {
	for _, teamMember := range l.team {
		teamMember.Stop()
	}
}
