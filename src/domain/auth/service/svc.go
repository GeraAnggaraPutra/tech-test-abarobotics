package service

import (
	"abarobotics-test/src/handler/database"
)

type Service struct {
	db database.DB
}

func NewService(db database.DB) *Service {
	return &Service{
		db: db,
	}
}
