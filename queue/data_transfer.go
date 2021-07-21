package queue

type (
	// Transfer the data between the tasks.
	Transfer interface {
		// Set the dataKey to be stored in the Transfer.
		Set(dataKey, interface{}) Transfer
		// Get the value from the Transfer.
		Get(dataKey) interface{}
	}

	transfer struct {
		data map[dataKey]interface{}
	}

	dataKey uint
)

const (
	// DataTransfer dataKey used for loading the transfer from the context.
	DataTransfer dataKey = iota
	// PayloadKey to load from the transfer.
	PayloadKey
)

// NewDataTransfer creates a Transfer instance.
func NewDataTransfer() Transfer {
	return transfer{
		data: make(map[dataKey]interface{}),
	}
}

func (t transfer) Set(key dataKey, value interface{}) Transfer {
	t.data[key] = value

	return t
}

func (t transfer) Get(key dataKey) interface{} {
	return t.data[key]
}
