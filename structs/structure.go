package structs

import "reflect"
import "../errors"

type structure struct {
	isPtr     bool
	structure interface{}
}

//FieldFoundCallback is a function called
type FieldFoundCallback func(*field)

//Struct returns a structure struct from provided struct
func Struct(i interface{}) (*structure, error) {

	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
			return nil, errors.ErrKindNotSupported
		}
		return &structure{isPtr: true, structure: i}, nil
	} else if reflect.TypeOf(i).Kind() != reflect.Struct {
		return nil, errors.ErrKindNotSupported
	}

	return &structure{structure: i}, nil
}

//Raw returns the raw struct or struct pointer
func (s *structure) Raw() interface{} {
	return s.structure
}

//fieldByIndex returns a pointer to a field struct from provided struct and index
func (s *structure) FieldByIndex(index int) (*field, error) {

	if s.FieldCount() <= index || (s.FieldCount() == 0 && index == 0) {
		return nil, errors.ErrFieldNotFound
	}

	if !s.isPtr {

		f := reflect.TypeOf(s.structure).Field(index)
		if f.PkgPath != "" {
			return nil, errors.ErrUnexportedField
		}
		if f.Anonymous {
			return nil, errors.ErrAnonymousField
		}

		return &field{field: f, value: reflect.ValueOf(s.structure).Field(index)}, nil
	}

	f := reflect.TypeOf(s.structure).Elem().Field(index)
	if f.PkgPath != "" {
		return nil, errors.ErrUnexportedField
	}

	return &field{field: f, value: reflect.ValueOf(s.structure).Elem().Field(index)}, nil
}

//fieldByName returns a pointer to a field struct from provided struct and name
func (s *structure) FieldByName(name string) (*field, error) {

	if !s.isPtr {

		f, success := reflect.TypeOf(s.structure).FieldByName(name)
		if !success {
			return nil, errors.ErrFieldNotFound
		} else if f.PkgPath != "" {
			return nil, errors.ErrUnexportedField
		}

		return &field{field: f, value: reflect.ValueOf(s.structure).FieldByName(name)}, nil

	}

	f, success := reflect.TypeOf(s.structure).Elem().FieldByName(name)
	if !success {
		return nil, errors.ErrFieldNotFound
	} else if f.PkgPath != "" {
		return nil, errors.ErrUnexportedField
	}

	return &field{field: f, value: reflect.ValueOf(s.structure).Elem().FieldByName(name)}, nil

}

//Name returns the name of the structure
func (s *structure) Name() string {

	if !s.isPtr {
		return reflect.TypeOf(s.structure).Name()
	}

	return reflect.TypeOf(s.structure).Elem().Name()
}

//Name returns the name of the structure
func (s *structure) FieldCount() int {

	if !s.isPtr {
		return reflect.TypeOf(s.structure).NumField()
	}

	return reflect.TypeOf(s.structure).Elem().NumField()
}

//fields returns a slice of field structs
func (s *structure) Fields() (fields []*field, err error) {

	for i := 0; i < s.FieldCount(); i++ {
		if f, err := s.FieldByIndex(i); err == nil {
			fields = append(fields, f)
		} else {
			if err != errors.ErrUnexportedField && err != errors.ErrAnonymousField {
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

	for _, field := range fields {
		m[field.Name()] = field.Value()
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

		values = append(values, field.Value())
	}
	return values, nil
}
