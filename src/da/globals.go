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
	log = logger
	var err error
	db, err := connectToPostgreSql()
	if err != nil {
		log.Error().Err(err).Msg(onInitErrMsg)
		return errors.New(onInitErrMsg)
	}
	Db = &SimpleWrapperSqlDb{
		db: db,
	}
	err = initDb(db)
	if err != nil {
		log.Error().Err(err).Msg(onInitErrMsg)
		return errors.New(onInitErrMsg)
	}

	initEmployeeRepo()

	return nil
}

func initEmployeeRepo() {
	EmployeeRepo = &PsqlEmployeeRepository{}
}

func OnClose() {
	if Db != nil {
		Db.Close()
	}
}
