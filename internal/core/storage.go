package core

import (
	"database/sql"
	"encoding/json"
	"errors"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID   int
	PID  int
	Cmd  string
	Args []string
	Log  string
	Time string
}

// -------- 新增任务 --------
func AddTask(t Task) error {
	argsJSON, _ := json.Marshal(t.Args)
	_, err := db.Exec(`INSERT INTO tasks (pid, cmd, args, log, time) VALUES (?, ?, ?, ?, ?)`,
		t.PID, t.Cmd, string(argsJSON), t.Log, t.Time)
	return err
}

// -------- 列出任务 --------
func ListTasks() ([]Task, error) {
	rows, err := db.Query(`SELECT id, pid, cmd, args, log, time FROM tasks ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		var argsJSON string
		if err := rows.Scan(&t.ID, &t.PID, &t.Cmd, &argsJSON, &t.Log, &t.Time); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(argsJSON), &t.Args)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func GetTaskPID(id int) (int, error) {
	var pid int
	err := db.QueryRow(`SELECT pid FROM tasks WHERE id = ?`, id).Scan(&pid)
	if err == sql.ErrNoRows {
		return 0, errors.New("task not found")
	}
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func DeleteTask(id int) error {
	_, err := db.Exec(`DELETE FROM tasks WHERE id = ?`, id)
	return err
}

func GetTaskLogPath(id int) (string, error) {
	var path string
	err := db.QueryRow(`SELECT log FROM tasks WHERE id = ?`, id).Scan(&path)
	if err == sql.ErrNoRows {
		return "", errors.New("task not found")
	}
	if err != nil {
		return "", err
	}
	return path, nil
}
