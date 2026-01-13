package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
)

var Version string

func main() {
	var title = flag.String("title", "Notification", "Notification title")
	var message = flag.String("message", "", "Notification message")
	var icon = flag.String("icon", "", "Notification icon path (optional)")
	var showVersion = flag.Bool("version", false, "Show version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", getVersion())
		os.Exit(0)
	}

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

func getVersion() string {
	if Version != "" {
		return Version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Version != "(devel)" {
			return info.Main.Version
		}
		if v, ok := getVCSBuildVersion(info); ok {
			return v
		}
	}
	return "(unset)"
}

func getVCSBuildVersion(info *debug.BuildInfo) (string, bool) {
	var (
		revision string
		dirty    string
	)
	for _, v := range info.Settings {
		switch v.Key {
		case "vcs.revision":
			revision = v.Value
		case "vcs.modified":
			dirty = " (dirty)"
		}
	}
	if revision == "" {
		return "", false
	}
	return revision + dirty, true
}
