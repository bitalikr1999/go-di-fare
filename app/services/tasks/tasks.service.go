package tasks_service

import (
	dbentities "bitalikr1999/difare/app/db/entities"
	"bitalikr1999/difare/app/db/repositories"
	"context"
	"fmt"
	"time"
)

type TasksService struct {
	tasksRepository *repositories.TasksRepository
}

func Create(tasksRepository *repositories.TasksRepository) *TasksService {

	return &TasksService{
		tasksRepository: tasksRepository,
	}
}

func (s *TasksService) GetList() dbentities.TasksList {

	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()

	list := s.tasksRepository.FindMany(ctx)

	fmt.Println(list)

	return list

}
