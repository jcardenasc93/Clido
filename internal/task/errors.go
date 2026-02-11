package task

import "fmt"

type TaskNotFoundError struct {
	id string
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with id %s not found", e.id)
}
