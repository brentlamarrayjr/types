package types

type channelQueue struct {
	closed   bool
	buffer   int
	elements chan interface{}
}

//ChannelQueue returns pointer to channelQueue struct
func ChannelQueue(maxCapacity int) *channelQueue {

	return &channelQueue{false, maxCapacity, make(chan interface{}, maxCapacity)}
}

func (q *channelQueue) Enqueue(i interface{}) error {

	if q.closed {
		return ErrQueueClosed
	}

	if q.Size() >= q.MaxCapacity() {
		return ErrMaxElements
	}

	q.elements <- i
	return nil
}

func (q *channelQueue) Dequeue() (interface{}, error) {
	if element, success := <-q.elements; success {
		return element, nil
	}
	return nil, ErrNoElements
}

func (q *channelQueue) Peek() (interface{}, error) {

	return nil, ErrMethodNotSupported
}

func (q *channelQueue) Size() int {

	return len(q.elements)
}

func (q *channelQueue) MaxCapacity() int {

	return q.buffer
}

func (q *channelQueue) IsEmpty() bool {

	return len(q.elements) > 0
}

func (q *channelQueue) Empty() {

	q.elements = make(chan interface{}, q.MaxCapacity())
}
