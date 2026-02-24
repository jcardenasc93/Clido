package cli

import (
	"flag"
	"fmt"
	"jcardenasc93/clido/internal/taskManager"
	"os"
	"strings"
)

type CLI struct {
	tm *taskmanager.TaskManager
}

func NewCLI(tm *taskmanager.TaskManager) *CLI {
	return &CLI{tm: tm}
}

func (c *CLI) Run() {
	if len(os.Args) < 2 {
		c.detailHandler("", []string{})
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "create":
		c.createHandler(args)
	case "detail":
		c.detailHandler("", args)
	default:
		// Fallback: if it's not a known command, treat everything as a description for create
		c.createHandler(os.Args[1:])
	}
}

func (c *CLI) createHandler(args []string) {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	desc := fs.String("description", "", "Task description")
	fs.Parse(args)

	finalDesc := *desc
	if finalDesc == "" {
		finalDesc = strings.Join(fs.Args(), " ")
	}

	if finalDesc == "" {
		fmt.Println("Error: description cannot be empty")
		os.Exit(1)
	}

	if err := c.tm.CreateTask(finalDesc); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Task created successfully.")
	c.tm.Tasks[len(c.tm.Tasks)-1].Pprint()
}

func (c *CLI) detailHandler(id string, args []string) {
	fs := flag.NewFlagSet("detail", flag.ExitOnError)
	idFlag := fs.String("id", id, "Task ID")
	fs.Parse(args)

	tID := *idFlag
	if tID != "" {
		task, err := c.tm.GetTaskByID(tID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		task.Pprint()
		return
	}

	for _, task := range c.tm.Tasks {
		task.Pprint()
	}
	fmt.Printf("\n%d tasks found\n", len(c.tm.Tasks))
}
