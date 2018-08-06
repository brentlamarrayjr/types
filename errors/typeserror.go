package errors

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
