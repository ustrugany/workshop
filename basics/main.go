package main

import (
	"fmt"
	"reflect"

	"golang.org/x/text/language"

	// 4. Easy package import by name
	"example.com/basics/quote"
)

// 1. Module
// - to initialize module run
// go mod init example.com/basics

// 2. To download dependencies
// go mod vendor
// inspect go.mod

// 3. Package main is the application and needs to implement main() method
func main() {
	var qp quote.Provider
	rp := quote.NewHelloProvider([]language.Tag{
		language.Make("en"),
		language.Make("fr"),
		language.Make("pl"),
		language.Make("ar"),
		language.Make("ar"),
		language.Make("el"),
		language.Make("sw"),
	})
	qp = rp
	q := qp.Provide()
	fmt.Println(q.Content)

	c := reflect.TypeOf(rp)
	fmt.Println("is provider implementing interface", c.Implements(reflect.TypeOf((*quote.Provider)(nil)).Elem()))
}
