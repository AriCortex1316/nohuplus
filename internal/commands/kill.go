package commands

import (
	"errors"
	"fmt"
	"syscall"

	"github.io/nohuplus/internal/core"
)

func KillTask(id int) error {
	pid, err := core.GetTaskPID(id)
	if err != nil {
		return err
	}
	defer core.DeleteTask(id)

	if err := syscall.Kill(pid, 0); err != nil {
		if errors.Is(err, syscall.ESRCH) {
			// 进程不存在
			return fmt.Errorf("process %d not found", pid)
		}
		if errors.Is(err, syscall.EPERM) {
			// 存在但没权限
			return fmt.Errorf("process %d exists but no permission to kill", pid)
		}
		return err // 其他错误
	}

	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return fmt.Errorf("failed to kill process %d: %w", pid, err)
	}
	return nil
}
