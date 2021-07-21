package tasks

import (
	"context"
	"fmt"
	"github.com/ahmedkamals/hashify/queue"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldMapCorrectly(t *testing.T) {
	testCases := []struct {
		id              string
		expected        string
		isErrorExpected bool
	}{
		{
			id:              "Should not return response from the server",
			expected:        "",
			isErrorExpected: true,
		},
		{
			id:              "Should save response from the server",
			expected:        "response",
			isErrorExpected: false,
		},
	}

	tasksFactory := NewTasksFactory()
	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			t.Parallel()

			testServer := getTestServer(testCase.expected)
			defer testServer.Close()

			dataTransfer := queue.NewDataTransfer().
				Set(queue.PayloadKey, testServer.URL)

			ctx := context.WithValue(context.Background(), queue.DataTransfer, dataTransfer)

			reduceQueue := make(chan queue.ResultSet, 1)
			defer close(reduceQueue)

			err := tasksFactory.CreateMappingTask(testServer.Client(), reduceQueue).Run(ctx)

			if testCase.isErrorExpected {
				assertEqual(t, testCase.expected, err)

				return
			}

			output := <-reduceQueue

			assertEqual(t, nil, err)
			assertEqual(t, testCase.expected, output.Result)
		})
	}
}

func BenchmarkDispatch(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	testServer := getTestServer("Hello world!!!")
	defer testServer.Close()

	dataTransfer := queue.NewDataTransfer().
		Set(queue.PayloadKey, testServer.URL)

	ctx := context.WithValue(context.Background(), queue.DataTransfer, dataTransfer)

	reduceQueue := make(chan queue.ResultSet, 100000)
	defer close(reduceQueue)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			NewTasksFactory().CreateMappingTask(testServer.Client(), reduceQueue).Run(ctx)
		}
	})
}

func getTestServer(response string) *httptest.Server {
	testServer := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, response)
			},
		),
	)

	testServer.EnableHTTP2 = true
	testServer.StartTLS()

	return testServer
}
