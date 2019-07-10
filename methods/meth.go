package methods

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   rune
}

type secretAgent struct {
	person
	ltk bool
}

// method
func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, sa.ltk)
}

// Meth tbd
func Meth() {
	fmt.Println("\nfunc Meth()")

	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	sa1 := secretAgent{
		person: person{
			first: "mikita",
			last:  "gulag",
			age:   23,
		},
		ltk: false,
	}

	fmt.Println(sa)

	sa.speak()
	sa1.speak()
}
