package task

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Task struct {
	id          string
	description string
	isDone      bool
}

func (t *Task) Pprint() {
	divisor := strings.Repeat("=", 60)
	fmt.Println(divisor)
	fmt.Printf("ID: %s\n", t.id)
	fmt.Printf("Description: %s\n", t.description)
	fmt.Printf("Status: %v\n", t.getStatus())
	fmt.Println(divisor)
}

func (t *Task) ToggleStatus() {
	t.isDone = !(t.isDone)
}

func (t *Task) getStatus() string {
	if t.isDone == true {
		return "Done"
	}
	return "Pending"
}

func AddTask(description string) Task {
	return Task{
		id:          uuid.NewString(),
		description: description,
	}
}
