package closure

import (
	"fmt"
)

// WizdomClosure tbd
func WizdomClosure() {
	fmt.Println("\nfunc WizdomClosure()")

	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
