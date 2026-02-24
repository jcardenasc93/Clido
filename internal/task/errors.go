package task

import (
	"errors"
	"fmt"
)

type TaskNotFoundError struct {
	ID string
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with id %s not found", e.ID)
}

var TaskWithNoDescriptionErr = errors.New("Task with no description is not allowed")

type NoValidTaskStatusErr struct {
	Status string
}

func (e *NoValidTaskStatusErr) Error() string {
	return fmt.Sprintf("Status value: %s is not valid", e.Status)
}
