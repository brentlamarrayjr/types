package types

import (
	"fmt"
)

type typesErr struct {
	Number      int
	Description string
}

func (e typesErr) Error() string {
	return fmt.Sprintf("Types(%d): %s", e.Number, e.Description)
}

//ErrTypeNotSupported is thrown when reflect.Type is not supported by function
var ErrTypeNotSupported = typesErr{Number: 1, Description: "Type not supported."}

//ErrNonStructPointer is thrown when reflect.Kind is not reflect.Pointer
var ErrNonStructPointer = typesErr{Number: 2, Description: "Pointer does not point to struct."}

//ErrFieldNotFound is thrown when reflect.StructField is not found in struct
var ErrFieldNotFound = typesErr{Number: 3, Description: "Field not found."}

//ErrTagNotFound is thrown when reflect.StructTag is not found in reflect.StructField
var ErrTagNotFound = typesErr{Number: 4, Description: "Tag not found."}

//ErrNoElements is thrown when there are no elements present
var ErrNoElements = typesErr{Number: 5, Description: "No elements."}

//ErrMaxElements is thrown when data structure has reached maximum amount of elements
var ErrMaxElements = typesErr{Number: 6, Description: "Max elements."}

//ErrMethodNotSupported is thrown when method is not supported
var ErrMethodNotSupported = typesErr{Number: 7, Description: "Method is not supported."}

//ErrValueNotSet is thrown when reflect.Value could not be set
var ErrValueNotSet = typesErr{Number: 8, Description: "Value not set."}

//ErrKindNotSupported is thrown when reflect.Kind is not supported by function
var ErrKindNotSupported = typesErr{Number: 9, Description: "Kind not supported."}

//ErrUnexportedField is thrown when reflect.Kind is not supported by function
var ErrUnexportedField = typesErr{Number: 10, Description: "Unexported field."}

//ErrQueueClosed is thrown when a type that implements queue has been closed
var ErrQueueClosed = typesErr{Number: 11, Description: "Queue is closed."}

//ErrAnonymousField is thrown when a reflect.StructField is anonymous
var ErrAnonymousField = typesErr{Number: 12, Description: "Field is anonymous."}
