package structs

import "reflect"
import "github.com/brentlamarrayjr/types/errors"

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

//GetType returns the struct type even for a struct pointer
func (s *structure) GetType() reflect.Type {
	if s.isPtr {
		return reflect.TypeOf(s.structure).Elem()
	}
	return reflect.TypeOf(s.structure)
}

//GetValue returns the struct value even for a struct pointer
func (s *structure) GetValue() reflect.Value {
	if s.isPtr {
		return reflect.ValueOf(s.structure).Elem()
	}
	return reflect.ValueOf(s.structure)
}

//FieldByIndex returns a pointer to a field struct from provided struct and index
func (s *structure) FieldByIndex(index int) (*field, error) {

	if index < 0 || s.FieldCount() == 0 || s.FieldCount() <= index {
		return nil, errors.ErrFieldNotFound
	}

	f := s.GetType().Field(index)
	if f.PkgPath != "" {
		return nil, errors.ErrUnexportedField
	}

	return &field{field: f, value: s.GetValue().Field(index), anonymous: f.Anonymous}, nil
}

//FieldByName returns a pointer to a field struct from provided struct and name
func (s *structure) FieldByName(name string) (*field, error) {

	f, success := s.GetType().FieldByName(name)
	if !success {
		return nil, errors.ErrFieldNotFound
	} else if f.PkgPath != "" {
		return nil, errors.ErrUnexportedField
	}

	return &field{field: f, value: s.GetValue().FieldByName(name), anonymous: f.Anonymous}, nil

}

//Name returns the name of the structure
func (s *structure) Name() string {
	return s.GetType().Name()
}

//FieldCount returns the count of fields in the supplied struct
//Unexported fields are skipped.
func (s *structure) FieldCount() int {
	return s.GetType().NumField()
}

func DeepFields(iface interface{}) []*field {

	fields := make([]*field, 0)

	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	if reflect.TypeOf(iface).Kind() == reflect.Ptr {

		ift = reflect.TypeOf(iface).Elem()
		ifv = reflect.ValueOf(iface).Elem()

	}

	for i := 0; i < ift.NumField(); i++ {

		t := ift.Field(i)
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			if t.Anonymous {
				fields = append(fields, DeepFields(v.Interface())...)
				continue
			}
			fallthrough
		case reflect.Ptr:
			if t.Anonymous && v.Elem().Kind() == reflect.Struct {
				fields = append(fields, DeepFields(v.Interface())...)
				continue
			}
			fallthrough
		default:
			fields = append(fields, &field{t, v, t.Anonymous})
		}
	}

	return fields
}

//Fields returns a slice of field structs.
//Embedded structs are flattened, including struct pointers.
//Unexported fields are skipped.
func (s *structure) DeepFields() ([]*field, error) {
	return DeepFields(s.structure), nil
}

//Fields returns a slice of field structs.
//Unexported fields are skipped.
func (s *structure) Fields() (fields []*field, err error) {

	for i := 0; i < s.FieldCount(); i++ {

		f, err := s.FieldByIndex(i)
		if err != nil {
			return nil, err
		}

		fields = append(fields, f)

	}

	return fields, nil
}

//Map returns a map of name and value pairs from fields.
//Unexported and anonymous fields are skipped.
func (s *structure) Map() (map[string]interface{}, error) {

	m := make(map[string]interface{}, s.FieldCount())

	fields, err := s.Fields()
	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		name, err := field.Name()
		if err != nil {
			if err == errors.ErrAnonymousField {
				continue
			}
			return nil, err
		}

		m[name] = field.Value()
	}

	return m, nil
}

//DeepMap returns a map of name and value pairs from fields.
//Embedded structs are flattened, including struct pointers.
//Unexported and anonymous fields are skipped.
func (s *structure) DeepMap() (map[string]interface{}, error) {

	m := make(map[string]interface{}, s.FieldCount())

	fields, err := s.DeepFields()
	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if !field.IsExported() || field.IsAnonymous() {
			continue
		}

		name, err := field.Name()
		if err != nil {
			return nil, err
		}

		m[name] = field.Value()
	}

	return m, nil
}

//Names returns a slice of names pull from fields.
//Unexported and anonymous fields are skipped.
func (s *structure) Names() (names []string, err error) {

	fields, err := s.Fields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if !field.IsExported() || field.IsAnonymous() {
			continue
		}

		name, err := field.Name()
		if err != nil {
			return nil, err
		}

		names = append(names, name)
	}
	return names, nil
}

//DeepNames returns a slice of names pull from fields.
//Embedded structs are flattened, including struct pointers.
//Unexported and anonymous fields are skipped.
func (s *structure) DeepNames() (names []string, err error) {

	fields, err := s.DeepFields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if !field.IsExported() || field.IsAnonymous() {
			continue
		}

		name, err := field.Name()
		if err != nil {
			return nil, err
		}

		names = append(names, name)
	}
	return names, nil
}

//DeepValues returns a slice of values pull from fields.
func (s *structure) Values() (values []interface{}, err error) {

	fields, err := s.Fields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if !field.IsExported() {
			continue
		}

		values = append(values, field.Value())
	}
	return values, nil
}

//DeepValues returns a slice of values pull from fields.
//Embedded structs are flattened, including struct pointers.
func (s *structure) DeepValues() (values []interface{}, err error) {

	fields, err := s.DeepFields()

	if err != nil {
		return nil, err
	}

	for _, field := range fields {

		if !field.IsExported() || field.IsAnonymous() {
			continue
		}

		values = append(values, field.Value())
	}
	return values, nil
}
