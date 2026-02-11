package main

import (
	"fmt"
	taskmanager "jcardenasc93/clido/internal/taskManager"
	"jcardenasc93/clido/storage"
)

func main() {
	jsonStorage := storage.NewJsonStorage("sample.json")
	taskManager := taskmanager.NewTaskManager(jsonStorage)
	for i := range 5 {
		description := fmt.Sprintf("A task #%d", i)
		err := taskManager.CreateTask(description)
		if err != nil {
			panic(fmt.Sprintf("Couldn't create task %d\n%v", i, err))
		}
	}
}
