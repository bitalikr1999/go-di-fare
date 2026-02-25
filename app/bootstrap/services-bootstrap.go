package bootstrap

import tasks_service "bitalikrty/difare/app/services/tasks"

type ServicesIoc struct {
	Tasks *tasks_service.TasksService
}
