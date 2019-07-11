package recursion

import (
	"fmt"
)

// GetRect tbd
func GetRect() {
	fmt.Println("\nfunc GetRect()")

	fmt.Println(factorial(4))
	fmt.Println(loopFuck(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// i'm the best
func loopFuck(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
