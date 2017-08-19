package tests

import (
	"fmt"
	"reflect"
	"testing"

	types "../../types"
	"github.com/stretchr/testify/require"
)

func TestMe(t *testing.T) {

	e := &Employee{}
	m := &Employee{ID: 1, Name: "manager", Manager: true, Salary: 50000, data: nil}
	fmt.Printf("Can set: %v \n", reflect.ValueOf(e).Elem().Field(0).CanSet())
	reflect.ValueOf(e).Elem().Field(0).Set(reflect.ValueOf(m.ID))
	fmt.Println(e.ID)

}

func TestFieldMethodName(t *testing.T) {

	e := &Employee{}

	s, err := types.Structure(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e).Elem().Kind())

	fields, err := s.Fields()
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	for i, field := range fields {
		fmt.Printf("(field) Name(%d): %s \n", i, field.Name())
	}
}

func TestFieldMethodSet(t *testing.T) {

	e := &Employee{}
	m := &Employee{ID: 1, Name: "manager", Manager: true, Salary: 50000, data: nil}

	s, err := types.Structure(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	s2, err := types.Structure(m)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(m))

	fields, err := s.Fields()
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	for i, field := range fields {

		f, err := s2.FieldByIndex(i)
		require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", i)
		require.Equalf(t, field.IsExported(), true, "true not returned via IsExported(%s) method of field", i)

		err = field.Set(f.Value())
		require.NoErrorf(t, err, "field struct could not be set via Set(%+v) method of field", f.Value())

	}

}
