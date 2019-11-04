package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"rsc.io/sampler"
)

func main() {
	// variable type will be derived from the value
	codes := []string{"en", "fr", "pl", "ar", "el", "sw"}

	// creating empty slice of type language.Tag
	var languageTags []language.Tag

	// slice size will be extended dynamically
	// while appending elements
	for _, l := range codes {
		languageTags = append(languageTags, language.Make(l))
	}

	// creating instance of Logger
	logger := logrus.New()
	for _, lt := range languageTags {
		logger.Infoln(sampler.Hello(lt))
	}
}
