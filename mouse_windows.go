// mouse_windows.go
//go:build windows
// +build windows

package main

import (
	"fmt"
	"syscall"
)

var (
	user32           = syscall.MustLoadDLL("user32.dll")
	procSetCursorPos = user32.MustFindProc("SetCursorPos")
)

func MoveMouse(x, y int) error {
	ret, _, err := procSetCursorPos.Call(uintptr(x), uintptr(y))
	if ret == 0 {
		return fmt.Errorf("failed to move mouse on Windows: %v", err)
	}
	return nil
}
