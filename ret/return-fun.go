package ret

import (
	"fmt"
)

// ReturnFunc tbd
func ReturnFunc() {
	fmt.Println("\nfunc ReturnFunc()")

	s := foo(42)
	fmt.Println(s)

	b := bar()()
	fmt.Println(b)
}

func foo(s rune) string {
	return fmt.Sprintf("just return argument %v", s)
}

func bar() func() int {
	return func() int {
		return 451
	}
}
