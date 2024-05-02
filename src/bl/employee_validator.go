package bl

import (
	"aiplus_demo/src/da"
	"fmt"
)

type EmployeeValidator interface {
	Validate(empl da.Employee) error
}

type ValidationError struct {
	Target string
	Reason string
}

func NewValidationError(target, reason string) error {
	return &ValidationError{}
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("failed to validate: %s, reason: %s", e.Target, e.Reason)
}
