package main

import (
	"fmt"
	"jcardenasc93/clido/internal/task"
)

func main() {
	var tasks []task.Task
	for i := range 5 {
		description := fmt.Sprintf("A task #%d", i)
		newTask := task.AddTask(description)
		tasks = append(tasks, newTask)
	}
	tasks[2].ToggleStatus()
	tasks[4].ToggleStatus()

	for _, t := range tasks {
		t.Pprint()
	}
}
