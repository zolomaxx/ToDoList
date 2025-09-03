package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

func main() {
	var tasks []Task
	nextID := 1
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Commands: add, list, complete, delete, exit")

	for {
		fmt.Print("\nEnter the command: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "add":
			fmt.Println("Enter a task: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			if title != "" {
				tasks = append(tasks, Task{ID: nextID, Title: title})
				fmt.Printf("Task added with ID %d\n", nextID)
				nextID++
			} else {
				fmt.Println("Error: task cannot be empty")
			}
		case "list":
			if len(tasks) == 0 {
				fmt.Println("The task list is empty")
			} else {
				fmt.Println("List of tasks:")
				for _, task := range tasks {
					status := " "
					if task.Completed {
						status = "✓"
					}
					fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
				}
			}
		case "complete":
			fmt.Print("Enter the task ID to complete: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Please enter a numeric ID")
				continue
			}

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					fmt.Printf("Task %d completed\n", id)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("Task with ID %d not found\n", id)
			}

		case "delete":
			fmt.Print("Enter the task ID to delete: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: Please enter a numeric ID")
				continue
			}

			found := false
			for i, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Task %d removed\n", id)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("Task with ID %d not found\n", id)
			}

		case "exit":
			fmt.Println("Exit the program")
			return

		case "":
		default:
			fmt.Println("Unknown command. Available commands: add, list, complete, delete, exit")
		}
	}
}
