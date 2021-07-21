package main

import (
	"context"
	"flag"
	"github.com/ahmedkamals/hashify/app"
	"github.com/ahmedkamals/hashify/config"
)

const defaultConcurrency uint = 10

func main() {
	config := parseFlags()
	args := flag.Args()

	config.QueueSize = uint(len(args))

	app.NewDispatcher(config).Dispatch(context.Background(), args)
}

func parseFlags() config.Configuration {
	concurrency := flag.Uint("concurrency", defaultConcurrency, "Max concurrent requests.")

	flag.Parse()

	return config.Configuration{
		Concurrency: *concurrency,
	}
}
