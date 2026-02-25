package bootstrap

import tasks_service "bitalikr1999/difare/app/services/tasks"

type ServicesIoc struct {
	Tasks *tasks_service.TasksService
}
