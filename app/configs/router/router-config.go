package config_router

import (
	tasks_service "bitalikrty/difare/app/services/tasks"
	"bitalikrty/difare/internal/router"
	"bitalikrty/difare/presentation/pages"
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
