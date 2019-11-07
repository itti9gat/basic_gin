package repo

import (
	"database/sql"
	"log"

	"iiujapp.tech/basic-gin/model"
)

// QueryData function
func QueryData(db *sql.DB) (model.ListUser, error) {

	var listData model.ListUser

	stmt, err := db.Prepare(`
		SELECT
			id, 
			username, 
			pwd, 
			name, 
			status 
		FROM users  
	`)
	if err != nil {
		return listData, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return listData, err
	}

	for rows.Next() {
		var user model.User

		if err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Name, &user.Status); err != nil {
			log.Fatal(err)
		}

		listData = append(listData, user)
	}

	return listData, nil
}

// WriteData function
func WriteData(db *sql.DB, m model.User) error {
	stmt, err := db.Prepare(`
		INSERT INTO users (
			id, 
			username, 
			pwd, 
			name, 
			status 
		) VALUES (
			NULL, ?, ?, ?, ?
		)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.Username,
		m.Password,
		m.Name,
		m.Status,
	)
	if err != nil {
		return err
	}

	return nil
}
