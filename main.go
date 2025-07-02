package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: todo [add|list|done] [args]")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		description := strings.Join(args[2:], " ")
		tasks, _ := LoadTasks()
		newTask := Task{
			ID:          len(tasks) + 1,
			Description: description,
			Completed:   false,
			CreatedAt:   time.Now(),
		}

		tasks = append(tasks, newTask)
		if err := SaveTasks(tasks); err != nil {
			fmt.Println("Failed to save task:", err)
			return
		}
		fmt.Println("Added task:", newTask.Description)

	case "list":
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, task := range tasks {
			PrintTask(task)
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide the task ID to delete.")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		newTasks := []Task{}
		found := false
		for _, task := range tasks {
			if task.ID == id {
				found = true
				continue // skip this task
			}
			newTasks = append(newTasks, task)
		}

		if !found {
			fmt.Printf("Task %d not found.\n", id)
			return
		}

		// Reassign new IDs to maintain order
		ReindexTasks(newTasks)

		if err := SaveTasks(newTasks); err != nil {
			fmt.Println("Failed to delete task:", err)
			return
		}

		fmt.Printf("Deleted task %d.\n", id)

	case "done":
		if len(args) < 3 {
			fmt.Println("Please provide the task ID to mark as done.")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		taskPtr, _ := FindTaskByID(tasks, id)
		if taskPtr == nil {
			fmt.Printf("Task %d not found.\n", id)
			return
		}

		taskPtr.Completed = true

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("Failed to update task:", err)
			return
		}
		fmt.Printf("Marked task %d as done.\n", id)

	case "edit":
		if len(args) < 4 {
			fmt.Println("Usage: todo edit <task_id> <new_description>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID.")
			return
		}

		newDescription := strings.Join(args[3:], " ")

		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		found := false
		for i, _ := range tasks {
			if tasks[i].ID == id {
				tasks[i].Description = newDescription
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Task %d not found.\n", id)
			return
		}

		if err := SaveTasks(tasks); err != nil {
			fmt.Println("Failed to update task:", err)
			return
		}

		fmt.Printf("Task %d updated successfully.\n", id)

	case "clear":
		tasks, err := LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}

		newTasks := []Task{}
		for _, task := range tasks {
			if !task.Completed {
				newTasks = append(newTasks, task)
			}
		}

		if len(newTasks) == len(tasks) {
			fmt.Println("No completed tasks to clear.")
			return
		}

		ReindexTasks(newTasks)

		if err := SaveTasks(newTasks); err != nil {
			fmt.Println("Failed to clear tasks:", err)
			return
		}

		fmt.Println("Cleared completed tasks.")

	case "help":
		fmt.Println("Usage: taskmanager <command> [arguments]")
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  add <description>        Add a new task")
		fmt.Println("  list                     List all tasks")
		fmt.Println("  done <task_id>           Mark a task as completed")
		fmt.Println("  edit <task_id> <text>    Edit a task's description")
		fmt.Println("  delete <task_id>         Delete a task")
		fmt.Println("  clear                    Delete all completed tasks")
		fmt.Println("  help                     Show this help message")

	default:
		fmt.Println("Unknown command:", command)
	}
}
