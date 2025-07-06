package model

import (
	"database/sql"
	"time"
)

type Device struct {
	GUID      string         `db:"guid"`
	Name      string         `db:"name"`
	Location  string         `db:"location"`
	Status    string         `db:"status"`
	CreatedAt time.Time      `db:"created_at"`
	CreatedBy sql.NullString `db:"created_by"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
	UpdatedBy sql.NullString `db:"updated_by"`
}
