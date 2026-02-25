package dbentities

import "time"

type Task struct {
	Id        int
	Title     string
	CreatedAt time.Time
}

type TasksList []Task
