package da

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const dbConnErrMsg = "failed to connect to psql db"

// todo add creation of table if it doesn't exist
func connectToPostgreSql() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv(dbUsernameEnv),
		os.Getenv(dbPasswordEnv),
		os.Getenv(dbNameEnv),
		os.Getenv(dbSslModeEnv),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Err(err).Msg(dbConnErrMsg)
		return nil, errors.New(dbConnErrMsg)
	}
	log.Info().Msg("connected to postgre")

	return db, nil
}

const initDbErr = "failed to init db"

func initDb(db *sql.DB) error {
	err := createEmployeesTableIfNeed(db)
	if err != nil {
		log.Error().Err(err).Msg(initDbErr)
		return errors.New(initDbErr)
	}
	return nil
}

const createTableErrTmplt = "failed to create table: %s"

func createEmployeesTableIfNeed(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS employees(
		id SERIAL PRIMARY KEY,
		firstName VARCHAR(20),
		middleName VARCHAR(20),
		lastName VARCHAR(20),
		phoneNumber VARCHAR(11) UNIQUE,
		city	VARCHAR(20)
		);`,
	)
	if err != nil {
		errMsg := fmt.Sprintf(createTableErrTmplt, "employees")
		log.Error().Err(err).Msg(errMsg)
		return errors.New(errMsg)
	}
	log.Info().Msg("initialized employee table")
	return nil
}
