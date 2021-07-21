package tasks

import (
	"context"
	"github.com/ahmedkamals/hashify/internal/errors"
	"github.com/ahmedkamals/hashify/queue"
	"io"
	"net/http"
)

// MappingTask applies a mapping function on a given input.
type MappingTask struct {
	queue.Task
	httpClient  *http.Client
	reduceQueue chan queue.ResultSet
}

// Run the MappingTask.
func (m *MappingTask) Run(controlCtx context.Context) error {
	const op errors.Operation = "MappingTask.Run"

	dataTransfer := controlCtx.Value(queue.DataTransfer).(queue.Transfer)
	url := dataTransfer.Get(queue.PayloadKey).(string)

	resp, err := m.httpClient.Get(url)

	if err != nil {
		return err
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.E(op, errors.Failure, err)
	}

	if err = resp.Body.Close(); err != nil {
		return errors.E(op, errors.Failure, err)
	}

	m.reduceQueue <- queue.ResultSet{
		URL:    url,
		Result: string(responseBody),
	}

	return nil
}
