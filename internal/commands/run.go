package commands

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"syscall"
	"time"

	"github.io/nohuplus/internal/core"
)

func RunCommand(cmd string, args []string) (*core.Task, error) {
	sum := md5.Sum([]byte(cmd))
	hash := hex.EncodeToString(sum[:])[:6]
	cmdName := filepath.Base(cmd)
	re := regexp.MustCompile(`[^a-zA-Z0-9._-]+`)
	cmdName = re.ReplaceAllString(cmdName, "_")
	startTime := time.Now()
	logPath := filepath.Join(core.AppPaths.LogDir,
		fmt.Sprintf("%s_%s_%s.log", hash, cmdName, startTime.Format("2006-01-02_150405")))
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	execCmd := exec.Command(cmd, args...)
	execCmd.Stdout = logFile
	execCmd.Stderr = logFile
	execCmd.Stdin = nil
	execCmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := execCmd.Start(); err != nil {
		return nil, err
	}

	finalLogPath := filepath.Join(core.AppPaths.LogDir,
		fmt.Sprintf("%s_%s_%s_%d.log", hash, cmdName, startTime.Format("2006-01-02_150405"), execCmd.Process.Pid))
	if err := os.Rename(logPath, finalLogPath); err == nil {
		logPath = finalLogPath
	}

	t := core.Task{
		PID:  execCmd.Process.Pid,
		Cmd:  cmd,
		Args: args,
		Log:  logPath,
		Time: time.Now().Format(time.RFC3339),
	}
	if err := core.AddTask(t); err != nil {
		return nil, err
	}

	return &t, nil
}
