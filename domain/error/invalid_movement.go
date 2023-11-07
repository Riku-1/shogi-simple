package error

import "fmt"

type InvalidMovementError struct {
	Msg string
}

func (e *InvalidMovementError) Error() string {
	return fmt.Sprintf("InvalidMovementError: %s", e.Msg)
}
