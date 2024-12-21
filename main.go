// main.go
package main

import (
	"fmt"
	m "github.com/snackbag/go-mousemover/mousemover"
	"os"
)

func main() {
	fmt.Println("Starting mouse movement test...")

	err := m.MoveMouse(20, 20)
	if err != nil {
		fmt.Println("Error moving mouse:", err)
		os.Exit(1)
	}
}
