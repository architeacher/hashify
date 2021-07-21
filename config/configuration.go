package config

type (
	// Configuration type wraps configuration data.
	Configuration struct {
		Concurrency, QueueSize uint
	}
)
