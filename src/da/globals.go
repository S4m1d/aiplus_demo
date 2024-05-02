package da

import (
	"errors"

	"github.com/rs/zerolog"
)

var Db WrapperSqlDb
var log zerolog.Logger
var EmployeeRepo SqlEmployeeRepository

const onInitErrMsg = "failed to init module da"

func OnInit(logger zerolog.Logger) error {
	logger = log
	var err error
	db, err := connectToPostgreSql()
	if err != nil {
		log.Error().Err(err).Msg(onInitErrMsg)
		return errors.New(onInitErrMsg)
	}
	Db = &SimpleWrapperSqlDb{
		db: db,
	}
	initEmployeeRepo()

	return nil
}

func initEmployeeRepo() {
	EmployeeRepo = &PsqlEmployeeRepository{}
}

func OnClose() {
	Db.Close()
}
