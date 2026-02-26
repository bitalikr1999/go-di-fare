package pages

import (
	events_tasks "bitalikr1999/difare/app/events/tasks"
	tasks_service "bitalikr1999/difare/app/services/tasks"
	eventsbus "bitalikr1999/difare/internal/events-bus"
	tasklist "bitalikr1999/difare/presentation/components/task-list"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type MainPage struct {
	tasksService *tasks_service.TasksService
	eventsBus    *eventsbus.EventsBus

	content fyne.CanvasObject
}

func CreateMain(
	tasksService *tasks_service.TasksService,
	eventsBus *eventsbus.EventsBus,
) *MainPage {
	page := &MainPage{
		tasksService: tasksService,
		eventsBus:    eventsBus,
	}
	go page.Listener()

	return page
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

	page.content = container.NewPadded(content)

	return page.content
}

func (page *MainPage) Refresh() {

	if page.content == nil {
		return
	}

	go fyne.Do(func() {
		page.content.Refresh()
		fmt.Println("Refreshed")
	})
}

func (page *MainPage) Listener() {

	chain := page.eventsBus.Subscribe(events_tasks.TASK_CREATED_EVENT)

	for event := range chain {
		taskCreatedEvent, ok := event.(events_tasks.TaskCreatedEventPayload)

		if !ok {
			fmt.Println("Error")
			continue
		}
		fmt.Println(taskCreatedEvent)
		page.Refresh()
	}
}

/*

When this page is loading it subscribes to some events
If events happen the page should be reloaded
*/
