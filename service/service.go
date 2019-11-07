package service

import (
	"database/sql"

	"iiujapp.tech/basic-gin/model"
)

// Service struct
type Service struct {
	DBData        *sql.DB
	QueryDataFunc queryDataFunc
	WriteDataFunc writeDataFunc
}

// Iservice interface
type Iservice interface {
	QueryUser() (model.ListUser, error)
	WriteData(m model.User) error
}

// QueryUser function
func (s Service) QueryUser() (model.ListUser, error) {
	return s.QueryDataFunc(s.DBData)
}

// WriteData function
func (s Service) WriteData(m model.User) error {
	return s.WriteDataFunc(s.DBData, m)
}
