package core

import (
	"os"
	"path/filepath"
)

type Paths struct {
	BaseDir string
	LogDir  string
}

var AppPaths *Paths

func EnsurePaths() error {
	// root 专用目录
	baseDir := "/var/lib/nohupuls"
	logDir := "/var/log/nohupuls"

	// 如果不是 root 用户，换到用户目录
	if os.Geteuid() != 0 {
		home, _ := os.UserHomeDir()
		baseDir = filepath.Join(home, ".local/share/nohupuls")
		logDir = filepath.Join(home, ".local/state/nohupuls/logs")
	}

	// 创建目录（如果不存在就自动新建）
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	AppPaths = &Paths{
		BaseDir: baseDir,
		LogDir:  logDir,
	}

	return nil
}
