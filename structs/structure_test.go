package structs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStructureInstantiation(t *testing.T) {

	e := &Employee{}

	_, err := Struct(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(e))

	s2 := []interface{}{0, "", false, 1.0}
	for _, element := range s2 {
		_, err := Struct(element)
		require.Errorf(t, err, "FAIL: structure struct instantiated via Struct(%s) method.", reflect.TypeOf(element))
	}

}

func TestStructureMethodFieldCount(t *testing.T) {

	m := &Manager{&Employee{0, "", 75000.50, nil}, true}

	s, err := Struct(m)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	count := s.FieldCount()
	require.IsTypef(t, 0, count, "FAIL: Did not return (%s) instantiated via FieldCount() method of structure", reflect.TypeOf(m))
	fmt.Printf("(structure) Count\n\tActual: %d\n\tExpected:%d", count, reflect.TypeOf(m).Elem().NumField())

}

func TestStructureMethodFieldByIndex(t *testing.T) {

	m := &Manager{&Employee{0, "", 75000.50, nil}, true}

	structure, err := Struct(m)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	_, err = structure.FieldByIndex(0)
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	_, err = structure.FieldByIndex(1)
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 1)

	_, err = structure.FieldByIndex(structure.FieldCount() + 1)
	require.Errorf(t, err, "field struct instantiated via FieldByIndex(%d) method of structure", structure.FieldCount()+1)

}

func TestStructureMethodFieldByName(t *testing.T) {

	m := &Manager{&Employee{0, "", 75000.50, nil}, true}

	structure, err := Struct(m)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	_, err = structure.FieldByName("ID")
	require.NoErrorf(t, err, "FAIL: field struct could not be instantiated via FieldByName(%s) method of structure", "ID")

	_, err = structure.FieldByName("data")
	require.Errorf(t, err, "FAIL: field struct could not be instantiated via FieldByName(%s) method of structure", "data")

	_, err = structure.FieldByName("Age")
	require.Errorf(t, err, "FAIL: field struct instantiated via FieldByName(%s) method of structure", "Age")

}

//
func TestStructureMethodFields(t *testing.T) {

	m := &Manager{&Employee{231345, "Brent", 75000.50, nil}, true}

	structure, err := Struct(m)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	fields, err := structure.Fields()
	require.NoError(t, err, "FAIL: *field slice not returned from Fields()")
	for _, field := range fields {
		fmt.Printf("(structure) Field: %+v \n", field)
	}

	fields, err = structure.DeepFields()
	require.NoError(t, err, "FAIL: *field slice not returned from DeepFields()")
	for _, field := range fields {
		fmt.Printf("(structure) Deep Field: %+v \n", field)
	}

}

func TestStructureMethodNames(t *testing.T) {

	m := &Manager{&Employee{231345, "Brent", 75000.50, nil}, true}

	structure, err := Struct(m)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	names, err := structure.Names()
	require.NoError(t, err, "FAIL: *field slice not returned from Names()")
	for _, name := range names {
		fmt.Printf("(structure) Name: %+v \n", name)
	}

	names, err = structure.DeepNames()
	require.NoError(t, err, "FAIL: *field slice not returned from DeepNames()")
	for _, name := range names {
		fmt.Printf("(structure) Deep Name: %+v \n", name)
	}

}

func TestStructureMethodMap(t *testing.T) {

	m := &Manager{&Employee{231345, "Brent", 75000.50, nil}, true}

	structure, err := Struct(m)
	require.NoErrorf(t, err, "FAIL: structure struct could not be instantiated via Struct(%s) method", reflect.TypeOf(m))

	sm, err := structure.Map()
	require.NoErrorf(t, err, "FAIL: map not returned via Map() method of structure. Error: ", err)
	fmt.Printf("Map: %+v", sm)

	dsm, err := structure.DeepMap()
	require.NoErrorf(t, err, "FAIL: map not returned via DeepMap() method of structure. Error: ", err)
	fmt.Printf("\nDeep Map: %+v", dsm)

}
