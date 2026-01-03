package commands

import (
	"errors"

	"github.io/nohuplus/internal/core"
)

func ListTasks() ([]core.Task, error) {
	list, err := core.ListTasks()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func ResolveTaskID(input int) (int, error) {
	list, err := core.ListTasks()
	if err != nil {
		return 0, err
	}
	if input <= 0 || input > len(list) {
		return 0, errors.New("task not found")
	}
	return list[input-1].ID, nil
}
