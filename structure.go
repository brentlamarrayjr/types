package types

import (
	"reflect"
)

var r = new(Reflection)

type structure struct {
	structure interface{}
}

//FieldFoundCallback is a function called
type FieldFoundCallback func(*field)

//Structure returns a structure struct from provided struct
func Structure(i interface{}) (*structure, error) {

	if reflect.TypeOf(i).Kind() != reflect.Ptr || reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
		return nil, ErrKindNotSupported
	}
	return &structure{structure: i}, nil
}

//fieldByIndex returns a pointet to a field struct from provided struct and index
func (s *structure) FieldByIndex(index int) (*field, error) {

	if s.FieldCount() <= index || (s.FieldCount() == 0 && index == 0) {
		return nil, ErrFieldNotFound
	}

	f := reflect.TypeOf(s.structure).Elem().Field(index)
	if f.PkgPath != "" {
		return nil, ErrUnexportedField
	}

	return &field{field: f, value: reflect.ValueOf(s.structure).Elem().Field(index)}, nil

}

//fieldByName returns a pointer to a field struct from provided struct and name
func (s *structure) FieldByName(name string) (*field, error) {

	f, success := reflect.TypeOf(s.structure).Elem().FieldByName(name)
	if !success {
		return nil, ErrFieldNotFound
	} else if f.PkgPath != "" {
		return nil, ErrUnexportedField
	}

	return &field{field: f, value: reflect.ValueOf(s.structure).Elem().FieldByName(name)}, nil

}

//Name returns the name of the structure
func (s *structure) Name() string {
	return reflect.TypeOf(s.structure).Elem().Name()
}

//Name returns the name of the structure
func (s *structure) FieldCount() int {
	return reflect.TypeOf(s.structure).Elem().NumField()
}

//fields returns a slice of field structs
func (s *structure) Fields() (fields []*field, err error) {

	for i := 0; i < s.FieldCount(); i++ {
		if f, err := s.FieldByIndex(i); err == nil {
			fields = append(fields, f)
		} else {
			if err != ErrUnexportedField {
				return nil, err
			}
		}
	}
	return fields, nil
}

//Map returns a map of fields represented by string/interface{} pairs
func (s *structure) Map() (map[string]interface{}, error) {

	m := make(map[string]interface{}, s.FieldCount())

	fields, err := s.Fields()
	if err != nil {
		return nil, err
	}

	counter := 0
	for _, field := range fields {

		if v, err := field.Value(); err == nil {
			m[field.Name()] = v
			counter++
		} else {
			return nil, err
		}
	}

	return m, nil
}

//Names returns a slice of field name strings
func (s *structure) Names() (names []string, err error) {

	fields, err := s.Fields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		names = append(names, field.Name())
	}
	return names, nil
}

//Values returns a slice of interface values from fields
func (s *structure) Values() (values []interface{}, err error) {

	fields, err := s.Fields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if v, err := field.Value(); err == nil {
			values = append(values, v)
		} else {
			return nil, err
		}

	}
	return values, nil
}
