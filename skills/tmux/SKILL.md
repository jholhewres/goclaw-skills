---
name: tmux
description: "Terminal multiplexing with tmux"
---
# tmux

Use the **bash** tool for tmux session management.

## Setup

1. **Check if installed:**
   ```bash
   command -v tmux && tmux -V
   ```

2. **Install:**
   ```bash
   # macOS
   brew install tmux

   # Ubuntu / Debian
   sudo apt update && sudo apt install -y tmux
   ```

## Sessions
```bash
tmux new -s <name>
tmux ls
tmux attach -t <name>
tmux kill-session -t <name>
tmux rename-session -t <old> <new>
```

## Windows & Panes
```bash
tmux new-window -t <session>
tmux split-window -h     # horizontal split
tmux split-window -v     # vertical split
tmux select-pane -t <n>
tmux list-windows -t <session>
```

## Send Commands
```bash
tmux send-keys -t <session>:<window> "command" Enter
tmux capture-pane -t <session>:<window> -p   # capture output
```

## Tips
- Use tmux send-keys to run commands in existing sessions
- Use tmux capture-pane -p to read session output
- Useful for managing long-running processes alongside the agent
