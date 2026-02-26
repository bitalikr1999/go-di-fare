package tasks_service

import (
	dbentities "bitalikr1999/difare/app/db/entities"
	"bitalikr1999/difare/app/db/repositories"
	events_tasks "bitalikr1999/difare/app/events/tasks"
	eventsbus "bitalikr1999/difare/internal/events-bus"
	"context"
	"fmt"
	"time"
)

type TasksService struct {
	tasksRepository *repositories.TasksRepository
	eventsBus       *eventsbus.EventsBus
}

func Create(tasksRepository *repositories.TasksRepository, eventsBus *eventsbus.EventsBus) *TasksService {

	return &TasksService{
		tasksRepository: tasksRepository,
		eventsBus:       eventsBus,
	}
}

func (s *TasksService) GetList() dbentities.TasksList {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	list := s.tasksRepository.FindMany(ctx)

	fmt.Println(list)

	return list
}

func (s *TasksService) Create() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.tasksRepository.Create(ctx, repositories.CreateTaskPayload{
		Title: "Some title",
	})

	fmt.Println("asdasd")

	s.eventsBus.Publish(events_tasks.TASK_CREATED_EVENT, events_tasks.TaskCreatedEventPayload{
		TaskId: 123,
	})
}
