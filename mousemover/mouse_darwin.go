// mouse_darwin.go
//go:build darwin
// +build darwin

package mousemover

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework ApplicationServices
#include <ApplicationServices/ApplicationServices.h>

int moveMouse(int x, int y) {
    CGEventRef moveEvent = CGEventCreateMouseEvent(NULL, kCGEventMouseMoved, CGPointMake(x, y), kCGMouseButtonLeft);
    if (moveEvent == NULL) {
        return -1;
    }
    CGEventPost(kCGHIDEventTap, moveEvent);
    CFRelease(moveEvent);
    return 0;
}
*/
import "C"
import "fmt"

func MoveMouse(x, y int) error {
	result := C.moveMouse(C.int(x), C.int(y))
	if result != 0 {
		return fmt.Errorf("failed to move mouse on macOS")
	}
	return nil
}
