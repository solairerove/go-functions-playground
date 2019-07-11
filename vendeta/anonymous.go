package vendeta

import (
	"fmt"
)

// Anonymous tbd
func Anonymous() {
	fmt.Println("\nfunc Anonymous()")

	foo()

	func() {
		fmt.Println("anonymous func()")
	}()

	func(x rune) {
		fmt.Println("anonymous func()", x)
	}(42)
}

func foo() {
	fmt.Println("func foo()")
}
