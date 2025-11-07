package commands

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.io/nohuplus/internal/core"
)

func ShowLog(id int) error {
	path, err := core.GetTaskLogPath(id)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}

func FollowLog(id int) error {
	path, err := core.GetTaskLogPath(id)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, _ := file.Stat()
	start := stat.Size()

	for {
		time.Sleep(500 * time.Millisecond)
		info, _ := file.Stat()
		if info.Size() < start {
			// 日志被清空或轮转
			start = 0
			file.Seek(0, io.SeekStart)
		} else if info.Size() > start {
			file.Seek(start, io.SeekStart)
			buf := make([]byte, info.Size()-start)
			n, _ := file.Read(buf)
			if n > 0 {
				os.Stdout.Write(buf[:n])
				start = info.Size()
			}
		}
	}
}
