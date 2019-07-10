package main

import (
	"fmt"

	def "github.com/solairerove/go-functions-playground/def"
	basic "github.com/solairerove/go-functions-playground/functions"
	meth "github.com/solairerove/go-functions-playground/methods"
	variadic "github.com/solairerove/go-functions-playground/variadic"
)

func main() {
	fmt.Println("go-functions-playground")

	basic.Basic()
	variadic.Variadic()
	def.DeferExample()
	meth.Meth()
}
