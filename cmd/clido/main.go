package main

import (
	"fmt"
	"jcardenasc93/clido/internal/cli"
	taskmanager "jcardenasc93/clido/internal/taskManager"
	"jcardenasc93/clido/storage"
	"os"
)

func main() {
	jsonStorage := storage.NewJsonStorage("tasks.json")
	taskManager := taskmanager.NewTaskManager(jsonStorage)

	if err := taskManager.LoadTasks(); err != nil {
		fmt.Printf("Couldn't load tasks from source: %v\n", err)
		os.Exit(1)
	}

	app := cli.NewCLI(taskManager)
	app.Run()
}
