package queues

//Queue is the representation of a basic queue
type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	Size() int
	Capacity() int
	MaxCapacity() int
	IsEmpty() bool
	IsMaxCapacity() bool
	Empty() error

	Close() error
}

//AsyncQueue is the representation of a basic queue supporting non blocking enqueue and dequeue
type AsyncQueue interface {
	EnqueueAsync(interface{}, func(error))
	DequeueAsync(func(interface{}, error))
	PeekAsync(func(interface{}, error))
	Size() int
	Capacity() int
	MaxCapacity() int
	IsEmpty() bool
	IsMaxCapacity() bool
	Empty() error

	Close() error
}
