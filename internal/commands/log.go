package commands

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.io/nohuplus/internal/core"
)

const tailLines = 10

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

	if err := printTail(file, tailLines); err != nil {
		return err
	}

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	start := stat.Size()

	for {
		time.Sleep(500 * time.Millisecond)
		info, err := file.Stat()
		if err != nil {
			return err
		}
		if info.Size() < start {
			// 日志被清空或轮转
			start = 0
			if _, err := file.Seek(0, io.SeekStart); err != nil {
				return err
			}
		} else if info.Size() > start {
			if _, err := file.Seek(start, io.SeekStart); err != nil {
				return err
			}
			buf := make([]byte, info.Size()-start)
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				return err
			}
			if n > 0 {
				if _, err := os.Stdout.Write(buf[:n]); err != nil {
					return err
				}
				start = info.Size()
			}
		}
	}
}

func printTail(file *os.File, n int) error {
	if n <= 0 {
		return nil
	}
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	size := stat.Size()
	if size == 0 {
		return nil
	}

	const chunkSize int64 = 4096
	var buf []byte
	var offset = size
	linesFound := 0

	for offset > 0 && linesFound <= n {
		readSize := chunkSize
		if offset < readSize {
			readSize = offset
		}
		offset -= readSize
		tmp := make([]byte, readSize)
		_, err := file.ReadAt(tmp, offset)
		if err != nil && err != io.EOF {
			return err
		}
		buf = append(tmp, buf...)
		linesFound += bytes.Count(tmp, []byte{'\n'})
	}

	parts := bytes.Split(buf, []byte{'\n'})
	if len(parts) > 0 && len(parts[len(parts)-1]) == 0 {
		parts = parts[:len(parts)-1]
	}
	if len(parts) > n {
		parts = parts[len(parts)-n:]
	}
	if len(parts) == 0 {
		return nil
	}

	_, err = os.Stdout.Write(bytes.Join(parts, []byte{'\n'}))
	return err
}
