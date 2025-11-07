# nohupuls

A lightweight personal task runner and logger — **not** a daemon, **not** a process supervisor.  
Just a small CLI tool that helps you run commands in the background, record their logs, and clean them up later.

---

## ✨ Features

- **No daemon needed** – nothing runs in the background except your own commands.
- **Structured task records** – each background task gets an entry with ID, PID, command, and log path.
- **Auto-generated logs** – each run produces a timestamped log file under `~/.local/state/nohupuls/logs`.
- **Kill and cleanup** – terminate tasks by ID or PID, and remove the record automatically.
- **Portable single binary** – build once, run anywhere (no dependencies or services required).

---

## 💡 Motivation

Sometimes you just need to start a script or binary in the background —  
nothing critical, nothing worth systemd, `tmux`, or a full-blown supervisor.

`nohup` works, but:
- It doesn’t remember what you’ve started.
- Logs are scattered or overwritten.
- Killing a process means hunting for its PID manually.

**nohupuls** fills that gap.  
It behaves like `nohup`, but keeps lightweight task metadata so you can:
- list what’s running,
- view logs quickly,
- kill and clean up with one command.

---

## ⚙️ Example

```bash
# Run a command
nohupuls run /home/user/1.sh

# Check all tasks
nohupuls list

# Kill a task by ID
nohupuls kill 3
