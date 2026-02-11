package storage

import (
	"jcardenasc93/clido/internal/task"
)

type Storage interface {
	SaveTasks([]*task.Task) error
	GetTaskById(string) (*task.Task, error)
	LoadTasks() ([]*task.Task, error)
}
