package da

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

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
		return nil, err
	}

	return db, nil
}
