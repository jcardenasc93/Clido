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
	if description == "" {
		return task.TaskWithNoDescriptionErr
	}
	t := task.CreateTask(description)
	m.Tasks = append(m.Tasks, t)
	err := m.storage.SaveTasks(m.Tasks)
	if err != nil {
		m.Tasks = m.Tasks[:len(m.Tasks)-1]
		return err
	}
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

func (m *TaskManager) GetTaskByID(id string) (*task.Task, error) {
	for _, t := range m.Tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return m.storage.GetTaskById(id)
}
