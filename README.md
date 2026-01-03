# 🌀 nohuplus

A lightweight personal task runner and logger — **not** a daemon, **not** a process supervisor.  
Just a small CLI tool that helps you run commands in the background, record their logs, and clean them up later.

---

## ✨ Features

- 💤 **No daemon needed** – nothing runs in the background except your own commands.  
- 🧾 **Persistent task store (SQLite)** – each background task is recorded with ID, PID, command, args, and log path.  
- 📜 **Auto-generated logs** – every run creates a timestamped log file under `~/.local/state/nohupuls/logs`.  
- 🧹 **Kill & cleanup** – terminate tasks by ID or PID, remove them from the database automatically.  
- 📦 **Portable single binary** – build once, run anywhere (no dependencies, no Python, no services).  

---

## 💡 Motivation

Sometimes you just need to start a script or binary in the background —  
nothing critical, nothing worth `systemd`, `tmux`, or a full-blown supervisor.

`nohup` works, but:

- It doesn’t remember what you’ve started.  
- Logs are scattered or overwritten.  
- Killing a process means hunting for its PID manually.  

**nohupuls** fills that gap.  
It behaves like `nohup`, but keeps lightweight task metadata in a tiny SQLite database so you can:

- list what’s running,  
- view logs quickly,  
- kill and clean up with one command.  

---

## ⚙️ Installation

```bash
wget https://github.com/AriCortex1316/nohuplus/releases/download/v0.2.0/nohuplus
chmod +x nohuplus
```
---

## 🧩 Usage

```bash
# Run a command in background
nohupuls bash /home/user/1.sh

# Check all tasks
nohupuls list

# View logs
nohupuls log 3          # show full log
nohupuls log -f 3       # follow (tail -f style)

# Kill a task by ID
nohupuls kill 3
```

---

## 🧠 Design Philosophy

- **Stateful nohup**: store metadata (PID, log path, args, timestamps) in SQLite.  
- **No daemon**: every action is executed on demand.  
- **Atomic updates**: SQLite ensures consistent task state.  
- **Simple over smart**: no watchdogs, no background service.  
- **Pure Go**: single binary, no runtime dependencies.

---

## 🪪 License

MIT © 2025  
Do whatever you want.
