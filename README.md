# noast

A cross-platform toast notification CLI application that works on Windows, Mac, WSL, and Linux.

**[日本語版 README](README_ja.md)**

## Installation

```bash
go install github.com/yashikota/noast@latest
```

## Usage

```bash
# Basic notification
noast -message "Hello, World!"

# Specify title and message
noast -title "Important" -message "Task completed"

# Specify icon
noast -title "Notification" -message "Message" -icon "/path/to/icon.png"
```
