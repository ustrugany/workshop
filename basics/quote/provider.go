// 5. Each folder is a package
package quote

import (
	"strings"

	"golang.org/x/text/language"
	"rsc.io/sampler"
)

type Quote struct {
	Content string
}

type Provider interface {
	Provide() Quote
}

type HelloProvider struct {
	// 6. Fields and methods starting from capital letter will be exported from package
	languageTags []language.Tag
}

// 8. Convention is to return pointer from constructor New* type of methods
func NewHelloProvider(tags []language.Tag) *HelloProvider {
	return &HelloProvider{
		languageTags: tags,
	}
}

// 7. Implicit interface implementation
func (hp *HelloProvider) Provide() Quote {
	var quotes []string
	for _, languageTag := range hp.languageTags {
		quotes = append(quotes, sampler.Hello(languageTag))
	}

	return Quote{Content: strings.Join(quotes, "\n")}
}
