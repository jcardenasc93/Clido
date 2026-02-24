package task

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t *Task) Pprint() {
	divisor := strings.Repeat("=", 60)
	fmt.Println(divisor)
	fmt.Printf("ID: %s\n", t.ID)
	fmt.Printf("Description: %s\n", t.Description)
	fmt.Printf("Status: %v\n", t.getStatus())
	fmt.Printf("Created: %v\n", formatTime(t.CreatedAt))
	fmt.Printf("Last Update: %v\n", formatTime(t.UpdatedAt))
	fmt.Println(divisor)
}

func (t *Task) ToggleStatus() {
	t.IsDone = !(t.IsDone)
}

func formatTime(time time.Time) string {
	// Layout string uses the reference date parts
	layout := "2006-01-02 15:04:05"
	return fmt.Sprintf("%s", time.Format(layout))

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
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *Task) Clone() *Task {
	return &Task{
		ID:          t.ID,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   time.Now(),
	}
}

var AllowedStatuses = []string{"done", "pending"}
