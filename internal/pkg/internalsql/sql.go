package internalsql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	return db, nil
}
