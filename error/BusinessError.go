package errors

import (
	"fmt"
)

type BusinessError struct {
	Reason string
}

func (e BusinessError) Error() string {
	return fmt.Sprintf("%v", e.Reason)
}
