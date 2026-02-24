package cli

import (
	"flag"
	"fmt"
	"jcardenasc93/clido/internal/task"
	"jcardenasc93/clido/internal/taskManager"
	"os"
	"slices"
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
	case "update":
		c.updateHandler("", args)
	case "filter":
		c.filterHandler(args)
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
	taskmanager.PPrintTasks(c.tm.Tasks, true)
}

func (c *CLI) updateHandler(id string, args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	idFlag := fs.String("id", id, "Task ID")
	descFlag := fs.String("description", "", "New description")
	statusFlag := fs.String("status", "", "Now task status")
	fs.Parse(args)

	tID := *idFlag
	if tID == "" {
		fmt.Printf("Error: Task ID is required to perform an update")
		os.Exit(1)
	}

	t, err := c.tm.UpdateTask(tID, *descFlag, *statusFlag)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}
	t.Pprint()
}

func (c *CLI) filterHandler(args []string) {
	fs := flag.NewFlagSet("filter", flag.ExitOnError)
	filterFlag := fs.String("type", "", "Filter type")
	valueFlag := fs.String("value", "", "Filter value")
	fs.Parse(args)

	filterType := *filterFlag
	filterValue := *valueFlag

	if filterType == "" {
		fmt.Printf("\nError: %v", task.MissingFilterTypeErr)
		os.Exit(1)
	}

	if filterValue == "" {
		fmt.Printf("\nError: %v", task.MissingFilterValueErr)
		os.Exit(1)
	}

	if slices.Contains(task.AllowedFilterTypes, filterType) == false {
		fmt.Printf("\nError: %v", task.NoValidFilterTypeErr{FilterType: filterType})
		os.Exit(1)
	}

	switch filterType {
	case "status":
		filter, err := task.NewStatusFilter(filterValue)
		if err != nil {
			fmt.Printf("\nError: %v", err)
			os.Exit(1)
		}
		tasks := filter.Filter(c.tm.Tasks)
		taskmanager.PPrintTasks(tasks, true)

	case "date-range":
		fmt.Println("Not implemented yet.")
	}
}
