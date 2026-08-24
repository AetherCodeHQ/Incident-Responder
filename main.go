package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: incident-responder <log> [ack|escalate|resolve]")
		os.Exit(1)
	}
	logPath := os.Args[1]
	action := "ack"
	if len(os.Args) > 2 {
		action = os.Args[2]
	}
	data, err := os.ReadFile(logPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")
	errors, warns := 0, 0
	for _, line := range lines {
		upper := strings.ToUpper(line)
		if strings.Contains(upper, "ERROR") || strings.Contains(upper, "FATAL") {
			errors++
		} else if strings.Contains(upper, "WARN") {
			warns++
		}
	}
	ts := time.Now().Format(time.RFC3339)
	switch action {
	case "ack":
		fmt.Printf("incident acknowledged at %s\n", ts)
		fmt.Printf("severity: %s\n", severity(errors, warns))
		fmt.Printf("errors: %d, warnings: %d\n", errors, warns)
	case "escalate":
		if errors > 5 {
			fmt.Println("ESCALATING to on-call (critical)")
		} else {
			fmt.Println("auto-handling (minor)")
		}
	case "resolve":
		fmt.Printf("resolved at %s - %d errors were reviewed\n", ts, errors)
	}
}

func severity(e, w int) string {
	if e > 10 {
		return "CRITICAL"
	} else if e > 0 {
		return "HIGH"
	} else if w > 0 {
		return "MEDIUM"
	}
	return "LOW"
}
