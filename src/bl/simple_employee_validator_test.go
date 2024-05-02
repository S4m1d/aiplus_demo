package bl

import (
	"aiplus_demo/src/da"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOk(t *testing.T) {
	re := regexp.MustCompile("^[0-9]{11}$")
	validator := &SimpleEmployeeValidator{
		re: re,
	}

	err := validator.Validate(da.Employee{
		FirstName:   "John",
		MiddleName:  "Johnovich",
		LastName:    "Doe",
		PhoneNumber: "88005553535",
		City:        "Bishkek",
	})

	assert.Nil(t, err)
}
