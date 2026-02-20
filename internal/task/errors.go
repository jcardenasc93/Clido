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
