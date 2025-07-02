package main

import (
	"encoding/json"
	"errors"
	"os"
)

const fileName = "task.json"

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644) // 0644 - file permission (owner - write, all - read)
}
