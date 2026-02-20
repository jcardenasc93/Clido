package storage

import (
	"encoding/json"
	"jcardenasc93/clido/internal/task"
	"log"
	"os"
	"path/filepath"
)

type JSONStorage struct {
	filePath string
}

func NewJsonStorage(filePath string) *JSONStorage {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := filepath.Join(home, ".local", "share", "clido")

	if err = os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatalf("Failed to create directory: %s\n%v", dirPath, err)
	}
	fullPath := filepath.Join(dirPath, filePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		if err := os.WriteFile(fullPath, []byte("[]"), 0644); err != nil {
			log.Fatalf("Failed to initialize storage file: %s\n%v", fullPath, err)
		}
	}

	return &JSONStorage{
		filePath: fullPath,
	}
}

func (s *JSONStorage) SaveTasks(t []*task.Task) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (s *JSONStorage) GetTaskById(id string) (*task.Task, error) {
	tasks, err := s.LoadTasks()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, &task.TaskNotFoundError{ID: id}
}

func (s *JSONStorage) LoadTasks() ([]*task.Task, error) {
	var tasks []*task.Task
	fileData, err := os.ReadFile(s.filePath)
	if err != nil {
		return tasks, &LoadFileError{filePath: s.filePath, details: err.Error()}
	}
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return tasks, &LoadFileError{filePath: s.filePath, details: err.Error()}
	}
	return tasks, nil
}
