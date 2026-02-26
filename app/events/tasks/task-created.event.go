package events_tasks

const TASK_CREATED_EVENT = "task.created"

type TaskCreatedEventPayload struct {
	TaskId int
}
