package query

import (
	"abarobotics-test/src/handler/database"

)

type Query struct {
	db database.Query
}

func NewQuery(db database.Query) *Query {
	return &Query{
		db: db,
	}
}
