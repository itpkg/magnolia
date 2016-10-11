package orm

import (
	"database/sql"
	"fmt"
	"os"
	"path"
)

var driver string
var db *sql.DB

func Open(drv, src string) error {
	var err error
	if db, err = sql.Open(drv, src); err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	driver = drv
	if err = os.MkdirAll(path.Join("db", drv, "migrations"), 0700); err != nil {
		return err
	}
	switch driver {
	case "postgres":
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS migration_schemes(id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL UNIQUE, created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW())")
	default:
		err = unknownDriverError

	}
	return err
}

var unknownDriverError = fmt.Errorf("unknown driver %s", driver)
