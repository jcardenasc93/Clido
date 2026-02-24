package taskmanager

import (
	"fmt"
	"jcardenasc93/clido/internal/task"
	"jcardenasc93/clido/storage"
	"slices"
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

func (m *TaskManager) UpdateTask(id string, desc string, status string) (*task.Task, error) {
	var tIdx int = -1
	var task_ *task.Task
	for i, t := range m.Tasks {
		if t.ID == id {
			task_ = t
			tIdx = i
			break
		}
	}
	if task_ == nil {
		return nil, &task.TaskNotFoundError{ID: id}
	}

	if desc == "" && status == "" {
		fmt.Println("Nothing to update.")
	}

	taskClone := task_.Clone()

	if desc != "" {
		taskClone.Description = desc
	}

	allowedStatus := []string{"done", "pending"}
	if status != "" {
		if slices.Contains(allowedStatus, status) == true {
			taskClone.IsDone = status == "done"
		} else {
			return nil, &task.NoValidTaskStatusErr{Status: status}
		}
	}

	// Create list of task to commit
	newTasks := make([]*task.Task, len(m.Tasks))
	copy(newTasks, m.Tasks)
	// Replace with the updated task
	newTasks[tIdx] = taskClone
	err := m.storage.SaveTasks(newTasks)
	if err != nil {
		return nil, err
	}
	// Update performed at storage level. Then manager slice must be updated
	m.Tasks[tIdx] = taskClone

	return taskClone, nil

}
