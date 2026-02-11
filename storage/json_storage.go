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
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirPath := filepath.Join(cwd, "exec", "output")

	if err = os.MkdirAll(dirPath, 0755); err != nil {
		log.Fatalf("Failed to create directory: %s\n%v", dirPath, err)
	}
	fullPath := filepath.Join(dirPath, filePath)
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

func (s *JSONStorage) DeleteTask(id string) error {
	return nil
}

func (s *JSONStorage) GetTaskById(id string) (*task.Task, error) {
	return nil, nil
}

func (s *JSONStorage) LoadTasks() []*task.Task {
	tasks := []*task.Task{}
	return tasks
}
