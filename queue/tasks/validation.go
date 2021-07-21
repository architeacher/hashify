package tasks

import (
	"context"
	"github.com/ahmedkamals/hashify/internal/errors"
	"github.com/ahmedkamals/hashify/queue"
	"regexp"
	"strings"
)

// ValidationTask validates input.
type ValidationTask struct {
	queue.Task
}

// Run the ValidationTask.
func (v *ValidationTask) Run(controlCtx context.Context) error {
	const op errors.Operation = "ValidationTask.Run"
	const urlPattern = `(?:https?:\/\/)?(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&//=]*)`
	const minURLLength = 3

	dataTransfer := controlCtx.Value(queue.DataTransfer).(queue.Transfer)
	url := dataTransfer.Get(queue.PayloadKey).(string)

	if len(url) < minURLLength {
		return errors.E(op, errors.MinLength)
	}

	re := regexp.MustCompile(urlPattern)

	if !re.MatchString(url) {
		return errors.E(op, errors.Invalid)
	}

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
		dataTransfer.Set(queue.PayloadKey, url)
	}

	return nil
}
