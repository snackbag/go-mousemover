// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting mouse movement test...")

	err := MoveMouse(20, 20)
	if err != nil {
		fmt.Println("Error moving mouse:", err)
		os.Exit(1)
	}
}
