package bootstrap

import (
	adapter_fyne "bitalikrty/difare/app/addapters/fyne"
	"bitalikrty/difare/app/configs/menu"
	config_router "bitalikrty/difare/app/configs/router"
	"bitalikrty/difare/app/db"
	"bitalikrty/difare/app/db/repositories"
	tasks_service "bitalikrty/difare/app/services/tasks"
	"bitalikrty/difare/internal/router"
	uilayouts "bitalikrty/difare/presentation/layouts"
	"context"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	fyneApp     fyne.App
	mainWindow  fyne.Window
	router      router.Router
	navigateCtr *router.RouterCtr

	repositories *RepositoriesIoc
	services     *ServicesIoc
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {

	a.initDb()
	a.initServices()
	a.initWindow()
	a.initRouter()

	go func() {
		time.Sleep(10 * time.Second)

		a.repositories.Tasks.Create(context.Background(), repositories.CreateTaskPayload{
			Title: "New task created dynamical",
		})

		fmt.Println("Created")
	}()

	a.mainWindow.ShowAndRun()

}

func (a *App) initWindow() {
	a.fyneApp = app.New()
	a.mainWindow = a.fyneApp.NewWindow("DiFare")

	a.mainWindow.Resize(fyne.NewSize(900, 500))
	a.mainWindow.CenterOnScreen()
}

func (a *App) initRouter() {

	a.navigateCtr = router.NewRouterCtr()

	layout := uilayouts.MainLayout{
		Items:       menu.NavMenuConfig,
		OnClickItem: a.navigateCtr.Navigate,
	}
	adapter := adapter_fyne.NewFyneRouterAdapter(a.mainWindow, &layout)

	a.router = *router.NewRouter(
		config_router.CreateRouterConfig(a.services.Tasks),
		config_router.Main,
		adapter,
		a.navigateCtr,
	)

	go a.router.Listen()
}

func (a *App) initDb() {

	ctx := context.Background()
	db, err := db.NewSQLite(ctx, "./db.sql")

	if err != nil {
		panic(err)
	}

	tasksRepository := repositories.NewTasksRepository(db)

	a.repositories = &RepositoriesIoc{
		Db:    db,
		Tasks: tasksRepository,
	}

}

func (a *App) initServices() {

	a.services = &ServicesIoc{
		Tasks: tasks_service.Create(a.repositories.Tasks),
	}

}
