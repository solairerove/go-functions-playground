package main

import (
	"fmt"
)

func main() {
	fmt.Println("go-functions-playground")

	foo()
	bar("James")

	w := woo("James")
	fmt.Println(w)

	s, b := gulag("gulag", "mikita")
	fmt.Println(s)
	fmt.Println(b)
}

func foo() {
	fmt.Println("hello from foo")
}

// pass by VALUE
func bar(s string) {
	fmt.Printf("hello from bar to %v\n", s)
}

func woo(s string) string {
	return fmt.Sprintf("hello from woo to %v\t", s)
}

func gulag(s1, s2 string) (string, bool) {
	s := fmt.Sprintf("hello from %v to %v\t", s1, s2)
	b := false

	return s, b
}
