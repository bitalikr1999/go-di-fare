package pages

import (
	tasks_service "bitalikr1999/difare/app/services/tasks"
	tasklist "bitalikr1999/difare/presentation/components/task-list"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainPage struct {
	tasksService *tasks_service.TasksService
}

func CreateMain(tasksService *tasks_service.TasksService) *MainPage {
	return &MainPage{
		tasksService: tasksService,
	}
}

func (page *MainPage) Render() fyne.CanvasObject {

	title := canvas.NewText("Tasks", color.White)

	tasks := page.tasksService.GetList()

	tasksListItems := []tasklist.Item{}

	for _, task := range tasks {

		tasksListItems = append(tasksListItems, tasklist.Item{
			Title: task.Title,
			Id:    task.Id,
		})

	}
	content := container.NewBorder(title, nil, nil, nil, tasklist.RenderList(tasksListItems))

	return container.NewPadded(content)
}
