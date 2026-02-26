package bootstrap

import (
	adapter_fyne "bitalikr1999/difare/app/addapters/fyne"
	"bitalikr1999/difare/app/configs/menu"
	config_router "bitalikr1999/difare/app/configs/router"
	"bitalikr1999/difare/app/db"
	"bitalikr1999/difare/app/db/repositories"
	tasks_service "bitalikr1999/difare/app/services/tasks"
	eventsbus "bitalikr1999/difare/internal/events-bus"
	"bitalikr1999/difare/internal/router"
	uilayouts "bitalikr1999/difare/presentation/layouts"
	"context"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type App struct {
	fyneApp      fyne.App
	mainWindow   fyne.Window
	router       router.Router
	navigateCtr  *router.RouterCtr
	repositories *RepositoriesIoc
	services     *ServicesIoc
	eventsBus    *eventsbus.EventsBus
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {

	a.initEventsBus()
	a.initDb()
	a.initServices()
	a.initWindow()
	a.initRouter()

	go func() {
		time.Sleep(1 * time.Second)

		a.services.Tasks.Create()

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
		config_router.CreateRouterConfig(a.services.Tasks, a.eventsBus),
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
		Tasks: tasks_service.Create(a.repositories.Tasks, a.eventsBus),
	}

}

func (a *App) initEventsBus() {
	a.eventsBus = eventsbus.New()
}
