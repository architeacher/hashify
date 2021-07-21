package app

import (
	"fmt"
	"github.com/ahmedkamals/hashify/config"
	"github.com/ahmedkamals/hashify/queue"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestReduce(t *testing.T) {
	testCases := []struct {
		id       string
		input    queue.ResultSet
		expected string
	}{
		{
			id: "Should encode the results",
			input: queue.ResultSet{
				Result: "abc",
			},
			expected: "900150983cd24fb0d6963f7d28e17f72",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			t.Parallel()

			reduceQueue := make(chan queue.ResultSet, 1)
			outputQueue := make(chan queue.ResultSet, 1)

			reduceQueue <- testCase.input

			go NewDispatcher(config.Configuration{}).reduce(reduceQueue, outputQueue, md5Encode)

			output := <-outputQueue
			assertEqual(t, testCase.expected, output.Result)
		})
	}
}

func TestDisplay(t *testing.T) {
	testCases := []struct {
		id       string
		input    queue.ResultSet
		expected string
	}{
		{
			id: "Should encode the results",
			input: queue.ResultSet{
				URL:    "http://www.hashify.com/",
				Result: "b4964faa7442e6ffbdda0bf08a6d50a9",
			},
			expected: fmt.Sprintf("%-28s %-32s \n", "http://www.hashify.com/", "b4964faa7442e6ffbdda0bf08a6d50a9"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			t.Parallel()

			outputQueue := make(chan queue.ResultSet, 1)
			outputQueue <- testCase.input
			close(outputQueue)

			output := captureOutput(t, func(output io.Writer) {
				NewDispatcher(config.Configuration{}).display(outputQueue)
				<-time.After(500 * time.Millisecond)
			})
			assertEqual(t, testCase.expected, output)
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

func captureOutput(t *testing.T, f func(output io.Writer)) string {
	t.Helper()

	rescueStdout := os.Stdout
	reader, writer, _ := os.Pipe()
	os.Stdout = writer

	f(writer)
	writer.Close()

	out, _ := ioutil.ReadAll(reader)
	os.Stdout = rescueStdout

	return string(out)
}
