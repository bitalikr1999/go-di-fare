package menu

import (
	config_router "bitalikr1999/difare/app/configs/router"
	navmenu "bitalikr1999/difare/presentation/components/nav-menu"
)

var NavMenuConfig = []navmenu.Item{
	{
		Label: "Main",
		Path:  config_router.Main,
	},
	{
		Label: "Settings",
		Path:  config_router.Settings,
	},
}
