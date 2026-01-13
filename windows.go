//go:build windows

package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func showWindowsNotification(title, message, icon string) error {
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

	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-Command", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}

func showNotification(title, message, icon string) error {
	return showWindowsNotification(title, message, icon)
}
