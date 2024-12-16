// mouse_linux.go
//go:build linux
// +build linux

package main

/*
#cgo LDFLAGS: -lX11
#include <X11/Xlib.h>

int moveMouse(int x, int y) {
    Display *display = XOpenDisplay(NULL);
    if (display == NULL) {
        return -1;
    }

    Window root = DefaultRootWindow(display);
    XWarpPointer(display, None, root, 0, 0, 0, 0, x, y);
    XFlush(display);
    XCloseDisplay(display);
    return 0;
}
*/
import "C"
import "fmt"

func MoveMouse(x, y int) error {
	result := C.moveMouse(C.int(x), C.int(y))
	if result != 0 {
		return fmt.Errorf("failed to move mouse on Linux")
	}
	return nil
}
