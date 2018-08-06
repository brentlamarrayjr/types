package queues

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"../errors"
)

func TestChannelQueue(t *testing.T) {

	capacity := 4

	q := ChannelQueue(capacity)

	require.Equalf(t, q.Size(), 0, "FAIL: element returned (%d) via Size() method does not match expected value (%d)", q.Size(), 0)
	require.Equalf(t, q.Capacity(), capacity, "FAIL: element returned (%d) via Capacity() method does not match expected value (%d)", q.Capacity(), capacity)
	require.True(t, q.IsEmpty(), "FAIL: queue should be empty")

	elements := []interface{}{0, "", 0.5, false}

	for _, element := range elements {
		err := q.Enqueue(element)
		require.NoErrorf(t, err, "FAIL: element could not be added to queue via Enqueue(%s) method", reflect.TypeOf(element))

	}

	require.False(t, q.IsEmpty(), "FAIL: queue should not be empty")

	for _, element := range elements {

		e, err := q.Dequeue()
		require.NoError(t, err, "FAIL: element not returned via Dequeue() method")
		require.Equalf(t, e, element, "FAIL: element returned (%v) via Dequeue() method does not match expected value (%v)", e, element)

	}

	_, err := q.Dequeue()
	require.Errorf(t, err, "FAIL: error (%s) not thrown via Dequeue() method when empty", errors.ErrNoElements)

	_, err = q.Peek()
	require.Errorf(t, err, "FAIL: error (%s) not thrown via Peek method when empty", errors.ErrMethodNotSupported)

}
