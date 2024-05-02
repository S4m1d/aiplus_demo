package bl

import (
	"aiplus_demo/src/da"
	"fmt"
	"regexp"
	"unicode"
)

type SimpleEmployeeValidator struct {
	re *regexp.Regexp
}

func (ev SimpleEmployeeValidator) Validate(empl da.Employee) error {
	if ok, validFieldDesc := isValidName(empl.FirstName); !ok {
		return NewValidationError("Employee.FirstName", validFieldDesc)
	}
	if ok, validFieldDesc := isValidName(empl.MiddleName); !ok {
		return NewValidationError("Employee.MiddleName", validFieldDesc)
	}
	if ok, validFieldDesc := isValidName(empl.LastName); !ok {
		return NewValidationError("Employee.LastName", validFieldDesc)
	}
	if ok, validFieldDesc := ev.isValidPhoneNumber(empl.PhoneNumber); !ok {
		return NewValidationError("Employee.PhoneNumber", validFieldDesc)
	}
	if ok, validFieldDesc := isValidName(empl.City); !ok {
		return NewValidationError("Employee.City", validFieldDesc)
	}

	return nil
}

const validNameDecription = "field must be non empty and start with capital letter"

func isValidName(name string) (bool, string) {
	return len(name) != 0 && unicode.IsUpper([]rune(name)[0]), validNameDecription
}

func (ev SimpleEmployeeValidator) isValidPhoneNumber(number string) (bool, string) {
	return ev.re.MatchString(number), fmt.Sprintf("field must satisfy the regexp pattern: %s", ev.re.String())
}
