package service

import (
	"database/sql"

	"iiujapp.tech/basic-gin/model"
)

type queryDataFunc func(db *sql.DB) (model.ListUser, error)

type writeDataFunc func(db *sql.DB, m model.User) error
