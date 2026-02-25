package bootstrap

import (
	"bitalikrty/difare/app/db/repositories"
	"database/sql"
)

type RepositoriesIoc struct {
	Db *sql.DB

	Tasks *repositories.TasksRepository
}
