package types

import "reflect"

type structure struct {
	isPtr     bool
	structure interface{}
}

//FieldFoundCallback is a function called
type FieldFoundCallback func(*field)

//Structure returns a structure struct from provided struct
func Structure(i interface{}) (*structure, error) {

	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		if reflect.TypeOf(i).Elem().Kind() != reflect.Struct {
			return nil, ErrKindNotSupported
		}
		return &structure{isPtr: true, structure: i}, nil
	} else if reflect.TypeOf(i).Kind() != reflect.Struct {
		return nil, ErrKindNotSupported
	}

	return &structure{structure: i}, nil
}

//fieldByIndex returns a pointet to a field struct from provided struct and index
func (s *structure) FieldByIndex(index int) (*field, error) {

	if s.FieldCount() <= index || (s.FieldCount() == 0 && index == 0) {
		return nil, ErrFieldNotFound
	}

	if !s.isPtr {

		f := reflect.TypeOf(s.structure).Field(index)
		if f.PkgPath != "" {
			return nil, ErrUnexportedField
		}
		if f.Anonymous {
			return nil, ErrAnonymousField
		}

		return &field{field: f, value: reflect.ValueOf(s.structure).Field(index)}, nil
	}

	f := reflect.TypeOf(s.structure).Elem().Field(index)
	if f.PkgPath != "" {
		return nil, ErrUnexportedField
	}

	return &field{field: f, value: reflect.ValueOf(s.structure).Elem().Field(index)}, nil
}

//fieldByName returns a pointer to a field struct from provided struct and name
func (s *structure) FieldByName(name string) (*field, error) {

	if !s.isPtr {

		f, success := reflect.TypeOf(s.structure).FieldByName(name)
		if !success {
			return nil, ErrFieldNotFound
		} else if f.PkgPath != "" {
			return nil, ErrUnexportedField
		}

		return &field{field: f, value: reflect.ValueOf(s.structure).FieldByName(name)}, nil

	}

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
			if err != ErrUnexportedField && err != ErrAnonymousField {
				return nil, err
			}
		}
	}
	return fields, nil
}

//Map returns a map of fields represented by string/interface{} pairs
func (s *structure) Map(lcase bool) (map[string]interface{}, error) {

	m := make(map[string]interface{}, s.FieldCount())

	fields, err := s.Fields()
	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		m[field.Name(lcase)] = field.Value()
	}

	return m, nil
}

//Names returns a slice of field name strings
func (s *structure) Names(lcase bool) (names []string, err error) {

	fields, err := s.Fields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {
		names = append(names, field.Name(lcase))
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
