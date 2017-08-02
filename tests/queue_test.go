package tests

import (
	types "../../types"
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestChannelQueue(t *testing.T) {

	q := types.ChannelQueue(2)

	q.Enqueue

	s, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	count := s.FieldCount()
	require.IsTypef(t, 0, count, "FAIL: Did not return (%s) instantiated via FieldCount() method of structure", reflect.TypeOf(e))
	fmt.Printf("(structure) Count: %d \n", count)
	fmt.Printf("(structure) Count: %d \n", reflect.TypeOf(e).Elem().NumField())
	fmt.Printf("(structure) Count: %d \n", reflect.TypeOf(e).Elem().NumField())

}
