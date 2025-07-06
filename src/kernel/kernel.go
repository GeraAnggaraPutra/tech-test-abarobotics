package kernel

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Kernel struct {
	dbx *sqlx.DB
	db  *gorm.DB
}

func NewKernel(
	dbx *sqlx.DB,
	db *gorm.DB,
) *Kernel {
	return &Kernel{
		dbx: dbx,
		db:  db,
	}
}

func (k *Kernel) GetDB() *gorm.DB {
	return k.db
}

func (k *Kernel) GetDBX() *sqlx.DB {
	return k.dbx
}
