package bootstrap

import (
	"bitalikr1999/difare/app/db/repositories"
	"database/sql"
)

type RepositoriesIoc struct {
	Db *sql.DB

	Tasks *repositories.TasksRepository
}
