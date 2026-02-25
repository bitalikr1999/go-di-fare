package menu

import (
	config_router "bitalikrty/difare/app/configs/router"
	navmenu "bitalikrty/difare/presentation/components/nav-menu"
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
