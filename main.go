package main

import (
	"fmt"
	"os"
)

// incident_responder - Auto incident response
func incident_responder(path string) {
	fmt.Println("========================================")
	fmt.Println("  Incident-Responder")
	fmt.Println("  Auto incident response")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	incident_responder(path)
}
