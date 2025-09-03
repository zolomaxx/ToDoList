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
		fmt.Print("\nВведите команду: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "add":
			fmt.Println("Введите задачу: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			if title != "" {
				tasks = append(tasks, Task{ID: nextID, Title: title})
				fmt.Printf("Задача добавлена с ID %d\n", nextID)
				nextID++
			} else {
				fmt.Println("Ошибка: задача не может быть пустой")
			}
		case "list":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
			} else {
				fmt.Println("Список задач:")
				for _, task := range tasks {
					status := " "
					if task.Completed {
						status = "✓"
					}
					fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
				}
			}
		case "complete":
			fmt.Print("Введите ID задачи для завершения: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Ошибка: введите числовой ID")
				continue
			}

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					fmt.Printf("Задача %d завершена\n", id)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("Задача с ID %d не найдена\n", id)
			}

		case "delete":
			fmt.Print("Введите ID задачи для удаления: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Ошибка: введите числовой ID")
				continue
			}

			found := false
			for i, task := range tasks {
				if task.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Задача %d удалена\n", id)
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("Задача с ID %d не найдена\n", id)
			}

		case "exit":
			fmt.Println("Выход из программы")
			return

		case "":
		default:
			fmt.Println("Неизвестная команда. Доступные команды: add, list, complete, delete, exit")
		}
	}
}
