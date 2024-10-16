package main

import "fmt"

type Task struct {
	Title       string
	Description string
	Status      string
	Priority    string
}

func addTask(tasks []Task, newTask Task) []Task {
	tasks = append(tasks, newTask)
	return tasks
}

func (t *Task) updateTaskStatus(newStatus string) {
	t.Status = newStatus
}

func getTasksByStatus(tasks []Task, status string) []Task {
	var newTasks []Task
	for _, val := range tasks {
		if val.Status == status {
			newTasks = append(newTasks, val)
		}
	}
	return tasks
}
func getTasksByPriority(tasks []Task, priority string) []Task {
	var newTasks []Task
	for _, val := range tasks {
		if val.Priority == priority {
			newTasks = append(newTasks, val)
		}
	}
	return newTasks
}

func printTasks(tasks []Task) {
	for _, val := range tasks {
		fmt.Printf("Title:\"%s\", Description:\"%s\", Status:\"%s\", Priority:\"%s\"\n", val.Title, val.Description, val.Status, val.Priority)
	}
}
func main() {
	var tasks []Task
	tasks = addTask(tasks, Task{"Title", "Description of task", "Completed", "Important"})
	tasks = addTask(tasks, Task{"Title", "Description of task", "Not completed", "Priority of task"})
	fmt.Println(tasks)
	fmt.Println()

	tasks[0].updateTaskStatus("Not completed")
	fmt.Println(tasks)
	fmt.Println()

	getByStatus := getTasksByStatus(tasks, "Not completed")
	fmt.Println(getByStatus)
	fmt.Println()

	getByPriority := getTasksByPriority(tasks, "Important")
	fmt.Println(getByPriority)

	fmt.Println()
	printTasks(tasks)
}
