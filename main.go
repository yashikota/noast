package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	var title = flag.String("title", "Notification", "Notification title")
	var message = flag.String("message", "", "Notification message")
	var icon = flag.String("icon", "", "Notification icon path (optional)")

	flag.Parse()

	if *message == "" {
		fmt.Fprintf(os.Stderr, "Error: message is required\n")
		flag.Usage()
		os.Exit(1)
	}

	err := showNotification(*title, *message, *icon)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error showing notification: %v\n", err)
		os.Exit(1)
	}
}

func showNotification(title, message, icon string) error {
	switch runtime.GOOS {
	case "windows":
		return showWindowsNotification(title, message, icon)
	case "darwin":
		return showMacNotification(title, message, icon)
	case "linux":
		if isWSL() {
			return showWSLNotification(title, message, icon)
		}
		return showLinuxNotification(title, message, icon)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
