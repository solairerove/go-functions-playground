package main

import (
	"fmt"

	def "github.com/solairerove/go-functions-playground/def"
	basic "github.com/solairerove/go-functions-playground/functions"
	jsf "github.com/solairerove/go-functions-playground/interfaces"
	meth "github.com/solairerove/go-functions-playground/methods"
	variadic "github.com/solairerove/go-functions-playground/variadic"
	vendeta "github.com/solairerove/go-functions-playground/vendeta"
)

func main() {
	fmt.Println("go-functions-playground")

	basic.Basic()
	variadic.Variadic()
	def.DeferExample()
	meth.Meth()
	jsf.JavaServerFaces()
	vendeta.Anonymous()
}
