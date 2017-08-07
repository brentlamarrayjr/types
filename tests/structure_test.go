package tests

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/brentlrayjr/types"
	"github.com/stretchr/testify/require"
)

func TestStructureInstantiation(t *testing.T) {

	e := &Employee{}

	_, err := types.Structure(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	s2 := []interface{}{0, "", false, 1.0}
	for _, element := range s2 {
		_, err := types.Structure(element)
		require.Errorf(t, err, "FAIL: structure struct instantiated via Structure(%s) method.", reflect.TypeOf(element))
	}

}

func TestStructureMethodFieldCount(t *testing.T) {

	e := &Employee{}

	s, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	count := s.FieldCount()
	require.IsTypef(t, 0, count, "FAIL: Did not return (%s) instantiated via FieldCount() method of structure", reflect.TypeOf(e))
	fmt.Printf("(structure) Count: %d \n", count)
	fmt.Printf("(structure) Count: %d \n", reflect.TypeOf(e).Elem().NumField())
	fmt.Printf("(structure) Count: %d \n", reflect.TypeOf(e).Elem().NumField())

}

func TestStructureMethodFieldByIndex(t *testing.T) {

	e := &Employee{}

	structure, err := types.Structure(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	_, err = structure.FieldByIndex(0)
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	_, err = structure.FieldByIndex(structure.FieldCount() + 1)
	require.Errorf(t, err, "field struct instantiated via FieldByIndex(%d) method of structure", structure.FieldCount()+1)

}

func TestStructureMethodFieldByName(t *testing.T) {

	e := &Employee{}

	structure, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	_, err = structure.FieldByName("ID")
	require.NoErrorf(t, err, "FAIL: field struct could not be instantiated via FieldByName(%s) method of structure", "ID")

	_, err = structure.FieldByName("data")
	require.Errorf(t, err, "FAIL: field struct could not be instantiated via FieldByName(%s) method of structure", "data")

	_, err = structure.FieldByName("Age")
	require.Errorf(t, err, "FAIL: field struct instantiated via FieldByName(%s) method of structure", "Age")

}

//
func TestStructureMethodFields(t *testing.T) {

	e := &Employee{}

	structure, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	fields, err := structure.Fields()
	require.NoError(t, err, "FAIL: *field slice not returned from Fields()")
	for _, field := range fields {
		fmt.Printf("(structure) Field: %+v \n", field)
	}
}

func TestStructureMethodNames(t *testing.T) {

	e := &Employee{}

	structure, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	fields, err := structure.Names(false)
	require.NoError(t, err, "FAIL: *field slice not returned from Names()")
	for _, field := range fields {
		fmt.Printf("(structure) Name: %+v \n", field)
	}
}

func TestStructureMethodMap(t *testing.T) {

	e := &Employee{}

	structure, err := types.Structure(e)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	m, err := structure.Map()
	require.NoErrorf(t, err, "FAIL: map not returned via Map() method of structure. Error: ", err)
	fmt.Printf("Map: %+v", m)

}
