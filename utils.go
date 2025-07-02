package main

import (
	"fmt"
)

func PrintTask(task Task) {
	status := "[-]"
	if task.Completed {
		status = "[x]"
	}
	fmt.Printf("%d. %s %s (created: %s)\n",
		task.ID, status, task.Description,
		task.CreatedAt.Format("2003-01-03 14:24"))
}

func FindTaskByID(tasks []Task, id int) (*Task, int) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], i
		}
	}
	return nil, -1
}

func ReindexTasks(tasks []Task) {
	for i := range tasks {
		tasks[i].ID = i + 1
	}
}
