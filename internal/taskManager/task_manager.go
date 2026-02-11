package taskmanager

import (
	"jcardenasc93/clido/internal/task"
	"jcardenasc93/clido/storage"
)

type TaskManager struct {
	storage storage.Storage
	Tasks   []*task.Task
}

func NewTaskManager(s storage.Storage) *TaskManager {
	return &TaskManager{
		storage: s,
	}
}

func (m *TaskManager) CreateTask(description string) error {
	t := task.CreateTask(description)
	err := m.storage.SaveTasks(m.Tasks)
	if err != nil {
		m.Tasks = m.Tasks[:len(m.Tasks)-1]
		return err
	}
	m.Tasks = append(m.Tasks, t)
	return nil
}

func (m *TaskManager) LoadTasks() error {
	tasks, err := m.storage.LoadTasks()
	if err != nil {
		return err
	}
	m.Tasks = tasks
	return err
}
