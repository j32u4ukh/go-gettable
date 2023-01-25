package internal

import "fmt"

// function in internal package which can not imported by others.
func Func() {
	fmt.Println("Call from package internal.")
}
