package task

import "slices"

var AllowedFilterTypes = []string{"status", "date-range"}

type TaskFilter interface {
	Filter(tasks []*Task) []*Task
}

type StatusFilter struct {
	IsDone bool
}

func NewStatusFilter(value string) (*StatusFilter, error) {
	if slices.Contains(AllowedStatuses, value) == true {
		return &StatusFilter{
			IsDone: value == "done",
		}, nil
	}
	return nil, &NoValidTaskStatusErr{Status: value}
}

func (sf *StatusFilter) Filter(tasks []*Task) []*Task {
	var result []*Task
	for _, t := range tasks {
		if t.IsDone == sf.IsDone {
			result = append(result, t)
		}
	}
	return result
}
