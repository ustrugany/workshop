package main

import (
	"errors"
	"fmt"

	// importing of your package
	"example.com/basics/sitemap"
)

func main() {
	// create your dependencies
	generator := sitemap.NewFakeGenerator(&sitemap.FakeRepository{})
	storage := sitemap.NewFileStorage("./sitemap.json")
	// always handle errors
	s, err := generator.Generate()
	if err != nil {
		var e sitemap.StorageError
		if errors.As(err, e) {
			fmt.Println("I think i can handle that!")
		} else {
			fmt.Println("uuups!")
		}
	}

	// always handle errors
	err = storage.Store(s)
	if err != nil {
		var e sitemap.StorageError
		if errors.As(err, e) {
			fmt.Println("I think i can handle that!")
		} else {
			fmt.Println("uuups!")
		}
	}

	fmt.Println("stored!")
}
