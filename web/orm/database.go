package orm

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

var driver string
var db *sql.DB
var unknownDriverError = fmt.Errorf("unknown driver %s", driver)

//Open open database
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
	if err != nil {
		return err
	}

	mpf := path.Join("db", drv, "mappers.toml")
	if _, err = os.Stat(mpf); os.IsNotExist(err) {
		log.Printf("grenate file %s", mpf)
		fd, er := os.OpenFile(mpf, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if er != nil {
			return er
		}
		defer fd.Close()
		err = toml.NewEncoder(fd).Encode(queries)
	}
	if err != nil {
		return err
	}
	log.Printf("load mappers")
	_, err = toml.DecodeFile(mpf, &queries)
	return err
}

//Exec db.Exec
func Exec(name string, args ...interface{}) (sql.Result, error) {
	qry, ok := queries[name]
	if !ok {
		return nil, fmt.Errorf("not found query for name %s", name)
	}
	return db.Exec(qry, args...)
}

//Query db.Query
func Query(name string, args ...interface{}) (*sql.Rows, error) {
	qry, ok := queries[name]
	if !ok {
		return nil, fmt.Errorf("not found query for name %s", name)
	}
	return db.Query(qry, args...)
}

//Query db.QueryRow
func QueryRow(name string, args ...interface{}) (*sql.Row, error) {
	qry, ok := queries[name]
	if !ok {
		return nil, fmt.Errorf("not found query for name %s", name)
	}
	return db.QueryRow(qry, args...), nil
}
