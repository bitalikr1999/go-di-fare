package repositories

import (
	dbentities "bitalikrty/difare/app/db/entities"
	"context"
	"database/sql"
	"fmt"
)

type TasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) *TasksRepository {
	return &TasksRepository{
		db: db,
	}
}

type CreateTaskPayload struct {
	Title string
}

func (r *TasksRepository) Create(ctx context.Context, payload CreateTaskPayload) {

	query := `
		INSERT INTO tasks (title) 
		VALUES (?)
	`

	_, err := r.db.ExecContext(ctx, query, payload.Title)

	if err != nil {
		fmt.Println("Error creating tasks", err)
	}

}

func (r *TasksRepository) Find(ctx context.Context) {

	query := `
		SELECT id, title, created_at FROM tasks 
		LIMIT 1
	`

	row := r.db.QueryRowContext(ctx, query)

	var task dbentities.Task

	err := row.Scan(
		&task.Id,
		&task.Title,
		&task.CreatedAt,
	)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(task)

}

func (r *TasksRepository) FindMany(ctx context.Context) dbentities.TasksList {

	query := `
		SELECT id, title, created_at FROM tasks 
		LIMIT 100
	`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return dbentities.TasksList{}
	}

	list := dbentities.TasksList{}

	for rows.Next() {
		var task dbentities.Task

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.CreatedAt,
		); err != nil {
			continue
		}

		list = append(list, task)

	}

	return list

}
