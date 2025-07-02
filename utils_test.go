package main

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintTask(t *testing.T) {
	for i := 0; i < 5; i++ {
		task_fail := Task{ID: i, Description: "Task which is not completed", Completed: false, CreatedAt: time.Now()}
		PrintTask(task_fail)
		task_success := Task{ID: i, Description: "Task which is completed", Completed: true, CreatedAt: time.Now()}
		PrintTask(task_success)
	}
}

func TestFindTaskByID(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1"},
		{ID: 2, Description: "Task 2"},
		{ID: 3, Description: "Task 3"},
		{ID: 4, Description: "Task 4"},
	}
	for i := 0; i < 5; i++ {
		i := i // capture range variable
		t.Run(fmt.Sprintf("Iteration-%d", i), func(t *testing.T) {
			task, idx := FindTaskByID(tasks, 3)
			if task == nil || idx != 2 {
				t.Errorf("Expected to find task with ID 3 at index 2 but got %v at index %d", task, idx)
			}
			task, idx = FindTaskByID(tasks, 99)
			if task != nil || idx != -1 {
				t.Errorf("Expected to not find task with ID 99 but got %v at index %d", task, idx)
			}
		})
	}
}

func TestReindexTasks(t *testing.T) {
	tasks := []Task{
		{ID: 13, Description: "Task 13"},
		{ID: 14, Description: "Task 14"},
	}
	for i := 0; i < 5; i++ {
		ReindexTasks(tasks)
		if tasks[0].ID != 1 || tasks[1].ID != 2 {
			t.Errorf("Iteration %d: Expected tasks to be reindexed to 1 and 2 but got %d and %d", i, tasks[0].ID, tasks[1].ID)
		}
	}
}

func TestReindexEmpty(t *testing.T) {
	for i := 0; i < 5; i++ {
		var tasks []Task
		ReindexTasks(tasks)
		if len(tasks) != 0 {
			t.Errorf("Iteration %d: Expected empty slice after reindex, got %v", i, tasks)
		}
	}
}

func TestClearCompletedTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Completed: true},
		{ID: 2, Completed: false},
		{ID: 3, Completed: true},
	}
	for i := 0; i < 5; i++ {
		newTasks := []Task{}
		for _, task := range tasks {
			if !task.Completed {
				newTasks = append(newTasks, task)
			}
		}
		if len(newTasks) != 1 || newTasks[0].ID != 2 {
			t.Errorf("Iteration %d: Clear completed tasks failed, got %v", i, newTasks)
		}
	}
}

func TestLoadAndSaveTasks(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task A", Completed: false, CreatedAt: time.Now()},
		{ID: 2, Description: "Task B", Completed: true, CreatedAt: time.Now()},
		{ID: 3, Description: "Task C", Completed: false, CreatedAt: time.Now().Add(-time.Hour)},
	}

	for i := 0; i < 5; i++ {
		if err := SaveTasks(tasks); err != nil {
			t.Fatalf("Iteration %d: SaveTasks failed: %v", i, err)
		}

		loaded, err := LoadTasks()
		if err != nil {
			t.Fatalf("Iteration %d: LoadTasks failed: %v", i, err)
		}

		if len(loaded) != len(tasks) {
			t.Fatalf("Iteration %d: expected %d tasks, got %d", i, len(tasks), len(loaded))
		}

		for j := range tasks {
			if loaded[j].ID != tasks[j].ID ||
				loaded[j].Description != tasks[j].Description ||
				loaded[j].Completed != tasks[j].Completed ||
				!loaded[j].CreatedAt.Round(time.Second).Equal(tasks[j].CreatedAt.Round(time.Second)) {
				t.Fatalf("Iteration %d: task mismatch at index %d", i, j)
			}
		}

		if err := SaveTasks([]Task{}); err != nil {
			t.Fatalf("Iteration %d: SaveTasks with empty slice failed: %v", i, err)
		}

		empty, err := LoadTasks()
		if err != nil {
			t.Fatalf("Iteration %d: LoadTasks with empty slice failed: %v", i, err)
		}
		if len(empty) != 0 {
			t.Fatalf("Iteration %d: expected empty slice", i)
		}
	}
}
