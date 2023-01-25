package pkg

import (
	"fmt"

	"github.com/j32u4ukh/gogettable/pkg/utils"
)

// Function in pkg package at the top of pkg folder.
func Func() {
	fmt.Println("Call from package pkg.")
}

// Call utils package from pkg package
func UtilsFunc() {
	utils.Func()
}
