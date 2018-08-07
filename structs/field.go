package structs

import (
	"reflect"

	"github.com/brentlamarrayjr/types/errors"
)

type field struct {
	field reflect.StructField
	value reflect.Value
}

func (f *field) Name() string {
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
	return "", errors.ErrTagNotFound
}

func (f *field) IsExported() bool {

	return f.field.PkgPath == ""
}

func (f *field) Set(value interface{}) error {

	if !f.IsExported() {
		return errors.ErrUnexportedField
	}

	if !f.value.CanSet() {
		return errors.ErrMethodNotSupported
	}

	if val := reflect.ValueOf(value); !val.IsValid() {
		return errors.ErrMethodNotSupported
	} else if val.Type().Kind() == reflect.Interface && val.IsNil() {
		f.value.Set(reflect.Zero(reflect.TypeOf(f.Value())))
		return nil
	}

	if f.value.Kind() != reflect.ValueOf(value).Kind() {
		return errors.ErrKindNotSupported
	}

	switch value.(type) {

	case int:
		f.value.SetInt(int64(value.(int)))

	case float64, float32:
		f.value.SetFloat(value.(float64))

	case string:
		f.value.SetString(value.(string))

	case bool:
		f.value.SetBool(value.(bool))

	default:
		f.value.Set(reflect.ValueOf(value))

	}

	return nil
}

func (f *field) IsZero() bool {

	return reflect.DeepEqual(f.Value(), reflect.Zero(f.field.Type).Interface())
}
