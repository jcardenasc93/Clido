package taskmanager

import (
	"jcardenasc93/clido/internal/task"
	"jcardenasc93/clido/storage"
)

type TaskManager struct {
	storage storage.Storage
	tasks   []*task.Task
}

func NewTaskManager(s storage.Storage) *TaskManager {
	return &TaskManager{
		storage: s,
	}
}

func (m *TaskManager) CreateTask(description string) error {
	t := task.CreateTask(description)
	err := m.storage.SaveTasks(m.tasks)
	if err != nil {
		m.tasks = m.tasks[:len(m.tasks)-1]
		return err
	}
	m.tasks = append(m.tasks, t)
	return nil
}
