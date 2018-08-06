package structs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

//Employee is a representation of an Employee for testing purposes
type Employee struct {
	ID      int
	Name    string `test:"testing"`
	Manager bool
	Salary  float64
	data    interface{}
}

func TestMe(t *testing.T) {

	e := &Employee{}
	m := &Employee{ID: 1, Name: "manager", Manager: true, Salary: 50000, data: nil}
	fmt.Printf("Can set: %v \n", reflect.ValueOf(e).Elem().Field(0).CanSet())
	reflect.ValueOf(e).Elem().Field(0).Set(reflect.ValueOf(m.ID))
	fmt.Println(e.ID)

}

func TestFieldMethodName(t *testing.T) {

	e := &Employee{}

	s, err := Struct(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e).Elem().Kind())

	fields, err := s.Fields()
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	for i, field := range fields {
		fmt.Printf("(field) Name(%d): %s \n", i, field.Name())
	}
}

func TestFieldMethodTag(t *testing.T) {

	e := &Employee{}

	s, err := Struct(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e).Elem().Kind())

	field, err := s.FieldByName("Name")
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByName(%s) method of structure", "Name")

	tag, err := field.Tag("test")
	require.NoErrorf(t, err, "Struct field should have tag", "Name")
	fmt.Println("TAG: " + tag)

}

func TestFieldMethodSet(t *testing.T) {

	e := &Employee{}

	s, err := Struct(e)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(e))

	fields, err := s.Fields()
	require.NoErrorf(t, err, "field struct could not be instantiated via FieldByIndex(%d) method of structure", 0)

	err = fields[0].Set(1)
	require.NoErrorf(t, err, "field struct could not be set via Set(%d) method of field", 1)
	fmt.Printf("Value(%d): %v \n", 0, fields[0].Value())

	err = fields[1].Set("manager")
	require.NoErrorf(t, err, "field struct could not be set via Set(%s) method of field", "manager")
	fmt.Printf("Value(%d): %v \n", 1, fields[1].Value())

	err = fields[2].Set(true)
	require.NoErrorf(t, err, "field struct could not be set via Set(%b) method of field", true)
	fmt.Printf("Value(%d): %v \n", 2, fields[2].Value())

	err = fields[3].Set(50000.00)
	require.NoErrorf(t, err, "field struct could not be set via Set(%d) method of field", 50000.00)
	fmt.Printf("Value(%d): %v \n", 3, fields[3].Value())

}
