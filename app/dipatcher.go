package app

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/ahmedkamals/hashify/config"
	"github.com/ahmedkamals/hashify/queue"
	"github.com/ahmedkamals/hashify/queue/tasks"
	"io"
	"log"
	"net/http"
)

type (
	// Dispatcher of the application.
	Dispatcher struct {
		config config.Configuration
	}

	// ReduceFunc used to reduce the output.
	ReduceFunc func(queue.ResultSet) queue.ResultSet
)

// NewDispatcher creates a new Dispatcher instance.
func NewDispatcher(config config.Configuration) *Dispatcher {
	return &Dispatcher{
		config: config,
	}
}

// Dispatch operation.
func (d Dispatcher) Dispatch(controlCtx context.Context, urls []string) {
	jobsQueue := make(chan queue.Job)
	reduceQueue := make(chan queue.ResultSet, d.config.QueueSize)
	outputQueue := make(chan queue.ResultSet, d.config.QueueSize)
	errorsQueue := make(chan error, d.config.QueueSize)

	go d.prepareJobs(urls, jobsQueue, reduceQueue)
	go d.reduce(reduceQueue, outputQueue, md5Encode)
	go d.display(outputQueue)

	leader := queue.NewLeader(jobsQueue, errorsQueue, d.config.Concurrency)
	leader.Run(controlCtx)
	leader.Wait()

	close(errorsQueue)

	d.reportErrors(errorsQueue)
}

func (d Dispatcher) prepareJobs(urls []string, jobsQueue chan queue.Job, reduceQueue chan queue.ResultSet) {
	for _, url := range urls {
		jobsQueue <- queue.NewJob(url, d.buildWorkflow(reduceQueue))
	}
	close(jobsQueue)
}

func (d Dispatcher) buildWorkflow(reduceQueue chan queue.ResultSet) queue.Workflow {
	tasksFactory := tasks.NewTasksFactory()

	tasksList := []queue.Task{
		tasksFactory.CreateValidationTask(),
		tasksFactory.CreateMappingTask(&http.Client{}, reduceQueue),
	}

	return queue.NewWorkflow(tasksList...)
}

func (d Dispatcher) reduce(reduceQueue chan queue.ResultSet, outputQueue chan<- queue.ResultSet, reduceFunc ReduceFunc) {
	for resultSet := range reduceQueue {
		go func(resultSet queue.ResultSet) {
			outputQueue <- reduceFunc(resultSet)
		}(resultSet)
	}
	close(reduceQueue)
	close(outputQueue)
}

func (d Dispatcher) display(outputQueue chan queue.ResultSet) {
	for resultSet := range outputQueue {
		go func(resultSet queue.ResultSet) {
			fmt.Printf("%-28s %-32s \n", resultSet.URL, resultSet.Result)
		}(resultSet)
	}
}

func (d Dispatcher) reportErrors(errorsQueue chan error) {
	fmt.Print("\nErrors:\n=======\n\n")

	for err := range errorsQueue {
		log.Println(err)
	}
}

func md5Encode(set queue.ResultSet) queue.ResultSet {
	h := md5.New()
	io.WriteString(h, set.Result)
	set.Result = fmt.Sprintf("%x", h.Sum(nil))

	return set
}
