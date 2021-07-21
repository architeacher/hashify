package queue

import "context"

type (
	// Workflow to be executed for a job.
	Workflow interface {
		// Execute one or more Task.
		Execute(context.Context) error
	}

	workflow struct {
		tasks []Task
	}
)

// NewWorkflow creates a new Workflow instance.
func NewWorkflow(tasks ...Task) Workflow {
	return &workflow{
		tasks: tasks,
	}
}

func (w workflow) Execute(controlCtx context.Context) error {
	for _, task := range w.tasks {
		if err := task.Run(controlCtx); err != nil {
			return err
		}
	}

	return nil
}
