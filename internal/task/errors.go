package task

import (
	"errors"
	"fmt"
)

type TaskNotFoundError struct {
	ID string
}

func (e TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with id %s not found.", e.ID)
}

var TaskWithNoDescriptionErr = errors.New("Task with no description is not allowed.")

type NoValidTaskStatusErr struct {
	Status string
}

func (e NoValidTaskStatusErr) Error() string {
	return fmt.Sprintf("Status value: %s is not valid.", e.Status)
}

type NoValidFilterTypeErr struct {
	FilterType string
}

func (e NoValidFilterTypeErr) Error() string {
	return fmt.Sprintf("Unexpected filter type: %s is not valid.", e.FilterType)
}

var MissingFilterTypeErr = fmt.Errorf("Filter type must be provided. Allowed: %v", AllowedFilterTypes)
var MissingFilterValueErr = errors.New("Filter value must be provided.")
