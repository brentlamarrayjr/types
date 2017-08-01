package types

import (
	"reflect"
)

type field struct {
	field reflect.StructField
	value reflect.Value
}

func (f *field) Name() string {
	return f.field.Name
}

func (f *field) Value() (interface{}, error) {

	if !f.IsExported() {
		return nil, ErrUnexportedField
	}

	return f.value.Interface(), nil
}

func (f *field) Index() int {
	return f.field.Index[0]
}

func (f *field) Tag(key string) (string, error) {

	if value, success := f.field.Tag.Lookup(key); success {

		return value, nil

	}
	return "", ErrTagNotFound
}

func (f *field) IsExported() bool {

	return f.field.PkgPath == ""
}

func (f *field) Set(value interface{}) error {

	if !f.IsExported() {
		return ErrUnexportedField
	}

	if !f.value.CanSet() {
		return ErrMethodNotSupported
	}

	if f.value.Kind() == reflect.ValueOf(value).Kind() {
		f.value.Set(reflect.ValueOf(value))
		return nil
	}
	return ErrValueNotSet
}

func (f *field) IsZero() bool {

	v, _ := f.Value()
	return reflect.DeepEqual(v, reflect.Zero(f.field.Type).Interface())
}
