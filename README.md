# noast

A cross-platform toast notification CLI application that works on Windows, WSL, Mac and Linux.  
Windown/WSL/Mac/Linux どれからでも使えるトースト通知用CLIアプリケーションです。  

## Installation

```bash
go install github.com/yashikota/noast@latest
```

or [Releases](https://github.com/yashikota/noast/releases)  

## Usage

```bash
# Basic notification
noast -message "Hello, World!"

# Specify title and message
noast -title "Important" -message "Task completed"

# Specify icon
noast -title "Notification" -message "Message" -icon "/path/to/icon.png"
```
