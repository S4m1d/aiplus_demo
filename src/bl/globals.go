package bl

import (
	"aiplus_demo/src/da"
	"os"
	"regexp"

	"github.com/rs/zerolog"
)

var log zerolog.Logger
var EmplWriter EmployeeWriter

func OnInit(logger zerolog.Logger) error {
	log = logger
	initEmplWriter()
	return nil
}

func initEmplWriter() {
	re := regexp.MustCompile(os.Getenv(phoneNumberRegexpEnv))
	validator := &SimpleEmployeeValidator{
		re: re,
	}

	EmplWriter = &SimpleEmployeeWriter{
		validator: validator,
		repo:      da.EmployeeRepo,
		db:        da.Db,
	}
}
