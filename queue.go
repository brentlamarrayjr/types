package types

//Queue is the representation of a basic queue
type Queue interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	IsMaxCapacity() bool
	Empty() error

	Close() error
}
