package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func isWSL() bool {
	if data, err := os.ReadFile("/proc/version"); err == nil {
		content := string(data)
		return strings.Contains(strings.ToLower(content), "microsoft") ||
			strings.Contains(strings.ToLower(content), "wsl")
	}

	if os.Getenv("WSL_DISTRO_NAME") != "" {
		return true
	}

	return false
}

func showWSLNotification(title, message, icon string) error {
	script := fmt.Sprintf(`
Add-Type -AssemblyName System.Windows.Forms
$notification = New-Object System.Windows.Forms.NotifyIcon
$notification.Icon = [System.Drawing.SystemIcons]::Information
$notification.BalloonTipIcon = [System.Windows.Forms.ToolTipIcon]::Info
$notification.BalloonTipText = '%s'
$notification.BalloonTipTitle = '%s'
$notification.Visible = $true
$notification.ShowBalloonTip(5000)
Start-Sleep -Seconds 1
$notification.Dispose()
`, strings.ReplaceAll(message, "'", "''"), strings.ReplaceAll(title, "'", "''"))

	cmd := exec.Command("powershell.exe", "-WindowStyle", "Hidden", "-Command", script)
	return cmd.Run()
}

func showLinuxNotification(title, message, icon string) error {
	args := []string{title, message}

	if icon != "" {
		args = append([]string{"--icon=" + icon}, args...)
	}

	if _, err := exec.LookPath("notify-send"); err != nil {
		return fmt.Errorf("notify-send not available on this Linux system")
	}

	cmd := exec.Command("notify-send", args...)
	return cmd.Run()
}
