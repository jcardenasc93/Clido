package main

import (
	"fmt"
	taskmanager "jcardenasc93/clido/internal/taskManager"
	"jcardenasc93/clido/storage"
	"log/slog"
)

func main() {
	jsonStorage := storage.NewJsonStorage("sample.json")
	taskManager := taskmanager.NewTaskManager(jsonStorage)

	// Load tasks if any
	err := taskManager.LoadTasks()
	if err != nil {
		slog.Error(err.Error())
	}

	fmt.Println()

	tasksQty := len(taskManager.Tasks)
	msg := fmt.Sprintf("%d tasks found", tasksQty)
	slog.Info(msg)

	for _, t := range taskManager.Tasks {
		t.Pprint()
	}

	for i := range 5 {
		description := fmt.Sprintf("A task #%d", i)
		err := taskManager.CreateTask(description)
		if err != nil {
			panic(fmt.Sprintf("Couldn't create task %d\n%v", i, err))
		}
	}
}
