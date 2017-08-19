package types

import (
	"reflect"
	"strings"
)

type field struct {
	field reflect.StructField
	value reflect.Value
}

func (f *field) Name(lcase bool) string {

	if lcase {
		return strings.ToLower(f.field.Name[:1]) + f.field.Name[1:]
	}
	return f.field.Name
}

func (f *field) Value() interface{} {

	return f.value.Interface()
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

	return reflect.DeepEqual(f.Value(), reflect.Zero(f.field.Type).Interface())
}
