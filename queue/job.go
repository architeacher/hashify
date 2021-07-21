package queue

import "context"

type (
	// Payload of the Job.
	Payload interface{}

	// Task to be applied inside a Workflow.
	Task interface {
		// Run as fast as it could.
		Run(context.Context) error
	}

	// Job definition.
	Job interface {
		// Payload returns the job's Payload.
		Payload() Payload

		// Workflow returns the job's to be applied Workflow.
		Workflow() Workflow
	}

	job struct {
		payload  Payload
		workFlow Workflow
	}
)

// NewJob creates a new Job instance.
func NewJob(payload Payload, workflow Workflow) Job {
	return &job{
		payload:  payload,
		workFlow: workflow,
	}
}

func (j job) Payload() Payload {
	return j.payload
}

func (j job) Workflow() Workflow {
	return j.workFlow
}
