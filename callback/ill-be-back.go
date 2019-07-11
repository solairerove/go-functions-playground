package callback

import (
	"fmt"
)

// BeBack tbd
func BeBack() {
	fmt.Println("\nfunc BeBack()")

	xi := []int{1, 2, 3, 4, 5}
	s := sum(xi...)
	fmt.Println(s)

	e := even(sum, xi...)
	fmt.Println(e)
}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func even(f func(x ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
