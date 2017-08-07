package types

type channelQueue struct {
	closed   bool
	buffer   int
	elements chan interface{}
}

//ChannelQueue returns pointer to channelQueue struct
func ChannelQueue(capacity int) *channelQueue {

	return &channelQueue{false, capacity, make(chan interface{}, capacity)}
}

func (q *channelQueue) Enqueue(i interface{}) error {

	if q.closed {
		return ErrQueueClosed
	}

	if q.Size() >= q.Capacity() {
		return ErrMaxElements
	}

	q.elements <- i
	return nil
}

func (q *channelQueue) Dequeue() (interface{}, error) {

	if q.IsEmpty() {
		return nil, ErrNoElements
	}

	element, success := <-q.elements
	if !success {
		return nil, ErrQueueClosed
	}

	return element, nil
}

func (q *channelQueue) Peek() (interface{}, error) {

	return nil, ErrMethodNotSupported
}

func (q *channelQueue) Size() int {

	return len(q.elements)
}

func (q *channelQueue) Capacity() int {

	return q.buffer
}

func (q *channelQueue) IsEmpty() bool {

	return len(q.elements) == 0
}

func (q *channelQueue) Empty() error {

	if q.closed {
		return ErrQueueClosed
	}

	q.elements = make(chan interface{}, q.Capacity())
	return nil
}

func (q *channelQueue) Close() error {

	if q.closed {
		return ErrQueueClosed
	}

	close(q.elements)
	q.closed = true
	return nil
}
