package core

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// 兜底：如果你已有 AppPaths.BaseDir，就删掉这个函数用你自己的路径。
func defaultBaseDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".local", "share", "nohuplus")
}

func InitDB() error {
	// 1) 准备路径
	base := defaultBaseDir() // 有 AppPaths 时换成 AppPaths.BaseDir
	if err := os.MkdirAll(base, 0o755); err != nil {
		return fmt.Errorf("mkdir base dir: %w", err)
	}
	dbPath := filepath.Join(base, "nohup_tasks.db")

	// 2) 打开数据库
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	// 3) 基本参数：WAL + busy timeout，利于并发与短暂锁等待
	if _, err := db.Exec(`PRAGMA journal_mode=WAL;`); err != nil {
		return fmt.Errorf("set WAL: %w", err)
	}
	if _, err := db.Exec(`PRAGMA busy_timeout=5000;`); err != nil {
		return fmt.Errorf("set busy_timeout: %w", err)
	}

	// 4) 建表（幂等）
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id   INTEGER PRIMARY KEY AUTOINCREMENT,
		pid  INTEGER,
		cmd  TEXT,
		args TEXT,   -- 存 []string 的 JSON
		log  TEXT,
		time TEXT
	);`)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() {
	if db != nil {
		_ = db.Close()
	}
}
