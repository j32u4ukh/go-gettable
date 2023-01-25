package gogettable

import (
	"fmt"

	"github.com/j32u4ukh/gogettable/internal"
	"github.com/j32u4ukh/gogettable/pkg"
	"github.com/j32u4ukh/gogettable/pkg/utils"
)

func Func() {
	fmt.Println("Hello, gogettable!")
}

func PkgFunc() {
	pkg.Func()
}

func UtilsFunc() {
	utils.Func()
}

func InternalFunc() {
	internal.Func()
}
