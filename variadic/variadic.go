package variadic

import (
	"fmt"
)

// Variadic tbd
func Variadic() {
	fmt.Println("\nfunc Variadic()")

	foo(1, 2, 3, 4, 5)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
