package model

import (
	"database/sql"
	"time"
)

type User struct {
	GUID        string         `db:"guid"`
	Email       string         `db:"email"`
	Password    string         `db:"password"`
	RoleGUID    string         `db:"role_guid"`
	RoleName    sql.NullString `db:"role_name"`
	CreatedAt   time.Time      `db:"created_at"`
	CreatedBy   sql.NullString `db:"created_by"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	UpdatedBy   sql.NullString `db:"updated_by"`
	Permissions []byte         `db:"permissions"`
}
