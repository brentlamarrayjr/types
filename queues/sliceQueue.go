package queues

import (
	"sync"

	"github.com/brentlamarrayjr/types/errors"
)

type sliceQueue struct {
	closed   bool
	buffer   int
	lock     *sync.Mutex
	elements []interface{}
}

//SliceQueue returns pointer to sliceQueue struct
func SliceQueue(capacity int) Queue {

	return &sliceQueue{false, capacity, &sync.Mutex{}, make([]interface{}, 0)}
}

func (q *sliceQueue) Enqueue(i interface{}) error {

	q.lock.Lock()

	if q.closed {
		return errors.ErrQueueClosed
	}

	if q.IsMaxCapacity() {
		return errors.ErrMaxElements
	}

	q.elements = append(q.elements, i)

	q.lock.Unlock()

	return nil
}

func (q *sliceQueue) EnqueueAsync(i interface{}, callback func(error)) {

	go func() {

		callback(q.Enqueue(i))

	}()

}

func (q *sliceQueue) Dequeue() (interface{}, error) {

	q.lock.Lock()

	if q.IsEmpty() {
		return nil, errors.ErrNoElements
	}

	element, elements := q.elements[0], q.elements[1:]

	q.elements = elements

	q.lock.Unlock()

	return element, nil
}

func (q *sliceQueue) DequeueAsync(callback func(interface{}, error)) {

	go func() {
		callback(q.Dequeue())
	}()

}

func (q *sliceQueue) Peek() (interface{}, error) {

	q.lock.Lock()

	if q.IsEmpty() {
		return nil, errors.ErrNoElements
	}

	q.lock.Unlock()

	return q.elements[0], nil

}

func (q *sliceQueue) PeekAsync(callback func(interface{}, error)) {

	go func() {
		callback(q.Peek())
	}()

}

func (q *sliceQueue) Size() int {

	return len(q.elements)
}

func (q *sliceQueue) MaxCapacity() int {
	return q.buffer
}

func (q *sliceQueue) Capacity() int {
	return q.MaxCapacity() - q.Size()
}

func (q *sliceQueue) IsMaxCapacity() bool {

	return q.Size() == q.MaxCapacity()
}

func (q *sliceQueue) IsEmpty() bool {

	return q.Size() == 0
}

func (q *sliceQueue) Empty() error {

	if q.closed {
		return errors.ErrQueueClosed
	}

	q.elements = make([]interface{}, q.MaxCapacity())

	return nil
}

func (q *sliceQueue) Close() error {

	if q.closed {
		return errors.ErrQueueClosed
	}

	q.closed = true
	return nil

}
