package migrations

import (
	"database/sql"
	"fmt"
)

func MigrateTaskTable(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
    		title TEXT NOT NULL,
    		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
