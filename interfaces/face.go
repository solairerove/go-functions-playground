package interfaces

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

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("Hi from bar", h)

	// conversion and asserting
	switch h.(type) {
	case person:
		fmt.Println("bar person", h.(person).first)
	case secretAgent:
		fmt.Println("bar agent", h.(secretAgent).first, h.(secretAgent).ltk)
	}
}

// JavaServerFaces tbd
func JavaServerFaces() {
	fmt.Println("\nfunc JavaServerFaces()")

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
	fmt.Printf("%T\n", sa)

	fmt.Println(sa1)
	fmt.Printf("%T\n", sa1)

	sa.speak()
	sa1.speak()

	bar(sa)
	bar(sa1)
}
