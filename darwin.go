//go:build darwin

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func showMacNotification(title, message, icon string) error {
	script := fmt.Sprintf(`display notification "%s" with title "%s"`,
		strings.ReplaceAll(message, `"`, `\"`),
		strings.ReplaceAll(title, `"`, `\"`))

	if icon != "" {
		script = fmt.Sprintf(`display notification "%s" with title "%s" subtitle "%s"`,
			strings.ReplaceAll(message, `"`, `\"`),
			strings.ReplaceAll(title, `"`, `\"`),
			"📢")
	}

	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

func showNotification(title, message, icon string) error {
	return showMacNotification(title, message, icon)
}
