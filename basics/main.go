// - to initialize module run go mod init example.com/basics
package main

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/text/language"

	// - easy package import by name
	"example.com/basics/quote"
)

// - to download dependencies go mod vendor, inspect go.mod
// - package main is the application and needs to implement main() method
func main() {
	fmt.Println("\n---")
	fmt.Println("declaration vs initialization")
	fmt.Println("---")

	// - this is declaration without initialization will return struct wth default zero values of its attributes
	// quote.ZeroStruct {Name: Elements:[] Counter:0 Flag:false Channel:<nil>}
	var z quote.ZeroStruct
	fmt.Printf("%T %+v\n", z, z)
	// - appending to a slice
	z.Elements = append(z.Elements, "one")
	// - worth noting that slices are passed via values, and will not reflect changes from inside of function
	func(elements []string) {
		// - this addition created new copy of slice
		elements = append(elements, "two")
		fmt.Printf("%T %+v\n", elements, elements)
	}(z.Elements)
	// - no changes from inside function
	fmt.Printf("%T %+v\n", z, z)

	fmt.Println("\n---")
	fmt.Println("pointers")
	fmt.Println("---")
	// - pointer to interface
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
	// - type of variable q (Quote) is implied by initial value
	// - don't do this, always handle errors
	q, _ := qp.Provide()
	fmt.Println("greetings")
	fmt.Println(q)

	fmt.Println("\n---")
	fmt.Println("error handling")
	fmt.Println("---")
	// - we will bind error value into this variable
	var e quote.ProviderError
	np := quote.NewNestedProvider(quote.NewFileProvider("/non/existing/path"))
	q, err := np.Provide()
	if err != nil {
		// - this mechanism will unwrap chain of errors and try to find error matching passed type
		if errors.As(err, &e) {
			fmt.Printf("unwrapped known error we can inspect the reason [%s] \n", e.Reason)
		} else {
			log.Fatalf("unknown error [%v]", err)
		}
	}

	fmt.Println("\n---")
	fmt.Println("non-uniform slices of objects of different types")
	fmt.Println("---")
	// - slice of variable types of objects
	providers := []interface{}{
		quote.NewFileProvider("file"),
		quote.NewHelloProvider(
			[]language.Tag{
				language.Make("en"),
			},
		),
		quote.Quote{},
	}
	for _, provider := range providers {
		// - type assertion to find only objects implementing Provider interface
		if p, ok := provider.(quote.Provider); ok {
			// - muting error not recommended
			q, _ = p.Provide()
			fmt.Println(q)
		}
	}
}
