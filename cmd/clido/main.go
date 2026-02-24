package main

import (
	"flag"
	"fmt"
	"jcardenasc93/clido/internal/task"
	taskmanager "jcardenasc93/clido/internal/taskManager"
	"jcardenasc93/clido/storage"
	"os"
	"strings"
)

func main() {
	jsonStorage := storage.NewJsonStorage("tasks.json")
	taskManager := taskmanager.NewTaskManager(jsonStorage)

	// Load tasks if any
	err := taskManager.LoadTasks()
	if err != nil {
		fmt.Printf("Couldn't load task from source\n%v\n", err)
		os.Exit(1)
	}

	cmdHandler(taskManager)

}

func cmdHandler(tm *taskmanager.TaskManager) {
	allowedCmd := []string{"create", "detail", "update", "delete"}
	if len(os.Args) < 1 {
		fmt.Printf("Missing command. Allowed ones: %v\n", allowedCmd)
		os.Exit(1)
	}

	subArgs := os.Args[2:]

	switch os.Args[1] {
	case "create":
		createCmd := flag.NewFlagSet("create", flag.ExitOnError)
		desc := createCmd.String("description", "", "Task description")
		createCmd.Parse(subArgs)
		if *desc == "" {
			*desc = strings.Join(createCmd.Args(), " ")

		}
		err := createTaskHandler(*desc, tm)
		if err != nil {
			fmt.Println("Couldn't create new task")
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println("New task created")
		t := tm.Tasks[len(tm.Tasks)-1]
		t.Pprint()

	case "detail":
		detailCmd := flag.NewFlagSet("detail", flag.ExitOnError)
		id := detailCmd.String("id", "", "Task UUID (optional)")
		detailCmd.Parse(subArgs)
		tasks, err := detailTaskHandler(*id, tm)
		if err != nil {
			fmt.Println("Something wrong while loading tasks")
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		for _, t := range tasks {
			t.Pprint()
		}
		tasksQty := len(tasks)
		msg := fmt.Sprintf("%d tasks found", tasksQty)
		fmt.Println(msg)

	default:
		fmt.Printf("Missing command. Allowed ones: %v\n", allowedCmd)
		os.Exit(1)
	}
}

func createTaskHandler(desc string, tm *taskmanager.TaskManager) error {
	return tm.CreateTask(desc)
}

func detailTaskHandler(id string, tm *taskmanager.TaskManager) ([]*task.Task, error) {
	tasks := []*task.Task{}
	if id != "" {
		t, err := tm.GetTaskByID(id)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, t)
		return tasks, nil
	}
	return tm.Tasks, nil
}
