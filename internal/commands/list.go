package commands

import (
	"github.io/nohuplus/internal/core"
)

func ListTasks() ([]core.Task, error) {
	list, err := core.ListTasks()
	if err != nil {
		return nil, err
	}

	return list, nil
}
