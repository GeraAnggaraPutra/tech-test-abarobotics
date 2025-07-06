package db

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

func openSQL(driver, source string, opt *databaseOption) (dbx *sqlx.DB, err error) {
	dbx, err = sqlx.Open(driver, source)
	if err != nil {
		return
	}

	dbx.SetMaxIdleConns(opt.connectionOption.maxIdle)
	dbx.SetMaxOpenConns(opt.connectionOption.maxOpen)
	dbx.SetConnMaxLifetime(opt.connectionOption.maxLifetime)
	dbx.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	return
}
