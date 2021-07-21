package tasks

import (
	"github.com/ahmedkamals/hashify/queue"
	"net/http"
)

type (
	// Factory a place where a Task is manufactured.
	Factory struct {
	}
)

// NewTasksFactory creates a new Factory instance.
func NewTasksFactory() *Factory {
	return new(Factory)
}

// CreateValidationTask creates a new instance of ValidationTask.
func (f Factory) CreateValidationTask() queue.Task {
	return new(ValidationTask)
}

// CreateMappingTask creates a new instance of MappingTask.
func (f Factory) CreateMappingTask(httpClient *http.Client, reduceQueue chan queue.ResultSet) queue.Task {
	return &MappingTask{
		httpClient:  httpClient,
		reduceQueue: reduceQueue,
	}
}
