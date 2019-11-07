package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"iiujapp.tech/basic-gin/conf"
)

// Start function
func Start() (*sql.DB, error) {

	configDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.MySqlUser, conf.MySqlPassword, conf.MySqlHost, conf.MySqlPort, conf.MySqlName)
	db, err := sql.Open("mysql", configDB)

	return db, err
}
