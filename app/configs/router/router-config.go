package config_router

import (
	tasks_service "bitalikr1999/difare/app/services/tasks"
	eventsbus "bitalikr1999/difare/internal/events-bus"
	"bitalikr1999/difare/internal/router"
	"bitalikr1999/difare/presentation/pages"
)

func CreateRouterConfig(
	tasksSevice *tasks_service.TasksService,
	eventsBus *eventsbus.EventsBus,
) router.ResourceList {

	return router.ResourceList{
		{
			Path:      Main,
			Component: pages.CreateMain(tasksSevice, eventsBus),
		},
		{
			Path:      Settings,
			Component: &pages.SettingsPage{},
		},
	}

}
