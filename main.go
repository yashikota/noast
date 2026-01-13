package main

import (
	"flag"
	"fmt"
	"os"
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
