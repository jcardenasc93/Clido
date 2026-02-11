package task

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func (t *Task) Pprint() {
	divisor := strings.Repeat("=", 60)
	fmt.Println(divisor)
	fmt.Printf("ID: %s\n", t.ID)
	fmt.Printf("Description: %s\n", t.Description)
	fmt.Printf("Status: %v\n", t.getStatus())
	fmt.Println(divisor)
}

func (t *Task) ToggleStatus() {
	t.IsDone = !(t.IsDone)
}

func (t *Task) getStatus() string {
	if t.IsDone == true {
		return "Done"
	}
	return "Pending"
}

func CreateTask(description string) *Task {
	id := uuid.NewString()
	return &Task{
		ID:          id,
		Description: description,
	}
}
