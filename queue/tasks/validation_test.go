package tasks

import (
	"context"
	"github.com/ahmedkamals/hashify/internal/errors"
	"github.com/ahmedkamals/hashify/queue"
	"reflect"
	"testing"
)

func TestValidation(t *testing.T) {
	testCases := []struct {
		id       string
		input    string
		expected error
	}{
		{
			id:       "Should return an error when url is too short",
			input:    "http://t.c",
			expected: errors.E(errors.Operation("ValidationTask.Run"), errors.MinLength),
		},
		{
			id:       "Should return an error when an invalid url",
			input:    "www.invalid-url,fyi",
			expected: errors.E(errors.Operation("ValidationTask.Run"), errors.Invalid),
		},
		{
			id:       "Should return no error when given a valid url without www",
			input:    "http://http://www.hashify.com/",
			expected: nil,
		},
		{
			id:       "Should return no error when given a valid url with www",
			input:    "http://www.twitter.com",
			expected: nil,
		},
	}

	tasksFactory := NewTasksFactory()
	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			t.Parallel()

			dataTransfer := queue.NewDataTransfer()
			dataTransfer.Set(queue.PayloadKey, testCase.input)

			ctx := context.WithValue(context.Background(), queue.DataTransfer, dataTransfer)
			err := tasksFactory.CreateValidationTask().Run(ctx)

			assertEqual(t, testCase.expected, err)
		})
	}
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	if (expected == nil || actual == nil) && expected != actual {
		reportError(t, expected, actual)
	}

	if !reflect.DeepEqual(expected, actual) {
		reportError(t, expected, actual)
	}
}

func reportError(t *testing.T, expected, actual interface{}) {
	t.Errorf("\nNot Equal:\nExpected: %v\nGot: %v", expected, actual)
}
