package gogettable

import (
	"fmt"

	"github.com/j32u4ukh/gogettable/v2/internal"
	"github.com/j32u4ukh/gogettable/v2/pkg"
	"github.com/j32u4ukh/gogettable/v2/pkg/utils"
)

// Function in gogettable package at the top of this project.
func Func() {
	fmt.Println("Hello, gogettable!")
}

// Call function in pkg package from gogettable package
func PkgFunc() {
	pkg.Func()
}

// Call function in utils package from gogettable package
func UtilsFunc() {
	utils.Func()
}

// Call function in internal package from gogettable package
func InternalFunc() {
	internal.Func()
}
