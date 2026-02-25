package config_router

import (
	tasks_service "bitalikr1999/difare/app/services/tasks"
	"bitalikr1999/difare/internal/router"
	"bitalikr1999/difare/presentation/pages"
)

func CreateRouterConfig(
	tasksSevice *tasks_service.TasksService,
) router.ResourceList {

	return router.ResourceList{
		{
			Path:      Main,
			Component: pages.CreateMain(tasksSevice),
		},
		{
			Path:      Settings,
			Component: &pages.SettingsPage{},
		},
	}

}
