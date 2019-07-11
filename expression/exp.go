package expression

import (
	"fmt"
)

// FuncExp tbd
func FuncExp() {
	fmt.Println("\nfunc FuncExp()")

	f := func() {
		fmt.Println("func expression")
	}
	f()

	bb := func(year rune) {
		fmt.Println("watching you since", year)
	}

	bb(1984)
}
