package queues

import (
	"../errors"
)

type channelQueue struct {
	closed   bool
	buffer   int
	elements chan interface{}
}

//ChannelQueue returns pointer to channelQueue struct
func ChannelQueue(capacity int) Queue {

	return &channelQueue{false, capacity, make(chan interface{}, capacity)}
}

func (q *channelQueue) Enqueue(i interface{}) error {

	if q.closed {
		return errors.ErrQueueClosed
	}

	if q.IsMaxCapacity() {
		return errors.ErrMaxElements
	}

	q.elements <- i
	return nil
}

func (q *channelQueue) EnqueueAsync(i interface{}, callback func(error)) {

	go func() {

		if q.closed {

			callback(errors.ErrQueueClosed)
			return
		}

		if q.IsMaxCapacity() {
			callback(errors.ErrMaxElements)
			return
		}

		q.elements <- i

		callback(nil)

	}()

}

func (q *channelQueue) Dequeue() (interface{}, error) {

	if q.IsEmpty() {
		return nil, errors.ErrNoElements
	}

	element, success := <-q.elements
	if !success {
		return nil, errors.ErrQueueClosed
	}

	return element, nil
}

func (q *channelQueue) DequeueAsync(callback func(interface{}, error)) {

	go func() {
		callback(q.Dequeue())
	}()

}

func (q *channelQueue) Peek() (interface{}, error) {

	return nil, errors.ErrMethodNotSupported
}

func (q *channelQueue) PeekAsync(callback func(interface{}, error)) {

	callback(q.Peek())
}

func (q *channelQueue) Size() int {

	return len(q.elements)
}

func (q *channelQueue) MaxCapacity() int {

	return q.buffer
}

func (q *channelQueue) Capacity() int {
	return q.MaxCapacity() - q.Size()
}

func (q *channelQueue) IsMaxCapacity() bool {

	return q.MaxCapacity() == q.Size()
}

func (q *channelQueue) IsEmpty() bool {

	return q.Size() == 0
}

func (q *channelQueue) Empty() error {

	if q.closed {
		return errors.ErrQueueClosed
	}

	q.elements = make(chan interface{}, q.Capacity())
	return nil
}

func (q *channelQueue) Close() error {

	if q.closed {
		return errors.ErrQueueClosed
	}

	close(q.elements)
	q.closed = true
	return nil
}
