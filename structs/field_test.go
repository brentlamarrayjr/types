package structs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

//Employee is a representation of an Employee for testing purposes
type Employee struct {
	ID     int
	Name   string `test:"testing"`
	Salary float64
	data   interface{}
}

type Manager struct {
	*Employee
	Override bool
}

func TestMe(t *testing.T) {

	e := &Employee{}
	m := &Employee{ID: 1, Name: "manager", Salary: 50000, data: nil}
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
		name, err := field.Name()
		require.NoErrorf(t, err, "Field at %d does not have a name", 0)
		fmt.Printf("(field) Name(%d): %s \n", i, name)
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

	m := &Manager{&Employee{}, true}

	s, err := Struct(m)
	require.NoErrorf(t, err, "structure struct could not be instantiated via Structure(%s) method", reflect.TypeOf(m))

	fields, err := s.Fields()
	require.NoError(t, err, "field struct could not be instantiated via Fields() method of structure")

	err = fields[1].Set(false)
	require.NoErrorf(t, err, "field struct could not be set via Set(%d) method of field", false)
	fmt.Printf("Value(%t): %t \n", false, fields[1].Value())

	err = fields[0].Set(&Employee{231345, "Brent", 75000.50, nil})
	require.NoErrorf(t, err, "field struct could not be set via Set(%b) method of field", true)
	fmt.Printf("Value(%+v): %+v \n", &Employee{231345, "Brent", 75000.50, nil}, fields[0].Value())

	fields, err = s.DeepFields()
	require.NoError(t, err, "field struct could not be instantiated via DeepFields() method of structure")

	err = fields[0].Set(12345)
	require.NoErrorf(t, err, "field struct could not be set via Set(%d) method of field", 12345)
	fmt.Printf("Value(%d): %v \n", 12345, fields[0].Value())

	err = fields[1].Set("manager")
	require.NoErrorf(t, err, "field struct could not be set via Set(%s) method of field", "manager")
	fmt.Printf("Value(%s): %v \n", "manager", fields[1].Value())

	err = fields[2].Set(50000.00)
	require.NoErrorf(t, err, "field struct could not be set via Set(%d) method of field", 50000.00)
	fmt.Printf("Value(%f): %f \n", 50000.00, fields[2].Value())

}
