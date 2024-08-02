package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: warn_script <url>")
		os.Exit(1)
	}

	url := os.Args[1]
	fmt.Printf("Running warn script for %s\n", url)

	// Add your warning script logic here
}