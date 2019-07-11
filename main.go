package main

import (
	"fmt"

	back "github.com/solairerove/go-functions-playground/callback"
	close "github.com/solairerove/go-functions-playground/closure"
	def "github.com/solairerove/go-functions-playground/def"
	exp "github.com/solairerove/go-functions-playground/expression"
	basic "github.com/solairerove/go-functions-playground/functions"
	jsf "github.com/solairerove/go-functions-playground/interfaces"
	meth "github.com/solairerove/go-functions-playground/methods"
	ret "github.com/solairerove/go-functions-playground/ret"
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
	exp.FuncExp()
	ret.ReturnFunc()
	back.BeBack()
	close.WizdomClosure()
}
