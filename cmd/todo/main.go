package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Revolyshn/todo-manager/internal/task"
	"github.com/Revolyshn/todo-manager/pkg/storage"
)

func main() {
	storage := storage.NewStorage("tasks.json")
	tasks, err := storage.Load()
	if err != nil {
		fmt.Println("Ошибка загрузки задач:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Добавить задачу")
		fmt.Println("2. Просмотреть задачи")
		fmt.Println("3. Отметить задачу как выполненную")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")
		fmt.Print("Выберите опцию: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addTask(reader, &tasks)
		case "2":
			viewTasks(tasks)
		case "3":
			completeTask(reader, &tasks)
		case "4":
			deleteTask(reader, &tasks)
		case "5":
			err := storage.Save(tasks)
			if err != nil {
				fmt.Println("Ошибка сохранения задач:", err)
			}
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверная опция, попробуйте снова.")
		}
	}
}

func addTask(reader *bufio.Reader, tasks *[]task.Task) {
	fmt.Print("Введите задачу: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	*tasks = append(*tasks, task.NewTask(description))
	fmt.Println("Задача добавлена.")
}

func viewTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println("Задач нет.")
		return
	}
	for i, task := range tasks {
		status := "Не выполнено"
		if task.Completed {
			status = "Выполнено"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
	}
}

func completeTask(reader *bufio.Reader, tasks *[]task.Task) {
	viewTasks(*tasks)
	fmt.Print("Введите номер задачи для отметки как выполненной: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index := parseIndex(indexStr)

	if index >= 0 && index < len(*tasks) {
		(*tasks)[index].Completed = true
		fmt.Println("Задача отмечена как выполненная.")
	} else {
		fmt.Println("Неверный номер задачи.")
	}
}

func deleteTask(reader *bufio.Reader, tasks *[]task.Task) {
	viewTasks(*tasks)
	fmt.Print("Введите номер задачи для удаления: ")
	indexStr, _ := reader.ReadString('\n')
	indexStr = strings.TrimSpace(indexStr)
	index := parseIndex(indexStr)

	if index >= 0 && index < len(*tasks) {
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
		fmt.Println("Задача удалена.")
	} else {
		fmt.Println("Неверный номер задачи.")
	}
}

func parseIndex(input string) int {
	var index int
	_, err := fmt.Sscanf(input, "%d", &index)
	if err != nil {
		return -1
	}
	return index - 1
}