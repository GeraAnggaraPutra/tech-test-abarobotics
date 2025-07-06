package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewDatabase() (dbx *sqlx.DB, db *gorm.DB, err error) {
	opt, err := newDatabaseOption()
	if err != nil {
		return
	}

	switch opt.driver {
	case "postgresql":
		dbx, db, err = NewPostgresql(opt)
	case "mysql":
		dbx, db, err = NewMysql(opt)
	case "":
	default:
		err = errors.Wrapf(errors.New("invalid datasources driver"), "db: driver=%s", opt.driver)
	}

	return
}
