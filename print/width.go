package print

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// getConsoleWidth calculates the width of the console/terminal window.
func getConsoleWidth() int {
	var size struct {
		Row, Col uint16
		X, Y     uint16
	}

	// Get the terminal size using syscall
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(os.Stdout.Fd()), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&size)))
	if err != 0 {
		fmt.Println("Error getting console size, defaulting to 80 characters wide.")
		return 80 // Default width if syscall fails
	}

	return int(size.Col)
}
