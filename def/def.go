package def

import (
	"fmt"
)

// DeferExample tbd
func DeferExample() {
	fmt.Println("\nfunc DeferExample()")

	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
