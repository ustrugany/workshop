// - each folder is a package
// - no classes, functionality is grouped in packages
package quote

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/language"
	"rsc.io/sampler"
)

// - this is what Go programmers talk about when they say “give your structs a useful zero value”
type ZeroStruct struct {
	Name string
	// - slices zero value is nil
	Elements []string
	Counter  int
	Flag     bool
	Channel  chan int
	// - pointer zero value is nil
	Pointer *ZeroStruct
}

// - data is grouped into structs
type Quote struct {
	Content string
}

// - interface
type Provider interface {
	Provide() (Quote, error)
}

type HelloProvider struct {
	// - fields and methods starting from capital letter will be exported from package, others are private
	languageTags []language.Tag
}

// - convention is to return pointer from constructor New* type of methods
func NewHelloProvider(tags []language.Tag) *HelloProvider {
	return &HelloProvider{
		languageTags: tags,
	}
}

// - implicit interface implementation
// - we define method Provide on struct HelloProvider
// - it is pointer type of method receiver
// - if function return error it should be last return argument
func (hp *HelloProvider) Provide() (Quote, error) {
	var quotes []string
	for _, languageTag := range hp.languageTags {
		quotes = append(quotes, sampler.Hello(languageTag))
	}

	return Quote{Content: strings.Join(quotes, "\n")}, nil
}

type FileProvider struct {
	path string
}

func NewFileProvider(path string) *FileProvider {
	return &FileProvider{path: path}
}

func (fp *FileProvider) Provide() (Quote, error) {
	var q Quote
	f, err := os.Open(fp.path)
	if err != nil {
		return q, ProviderError{Reason: fmt.Sprintf("file does not exists %q", fp.path), Err: err}
	}

	// - defer is a language mechanism that puts your function call into a stack, be called before returning from function
	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	var quotes []string
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return q, ProviderError{Reason: fmt.Sprintf("failed to scan file %q", fp.path), Err: err}
	}

	return Quote{Content: strings.Join(quotes, "\n")}, nil
}

type NestedProvider struct {
	provider Provider
}

// - we need to pass pointer to struct *FileProvider as pointer receiver implements the interface Provider
func NewNestedProvider(provider *FileProvider) *NestedProvider {
	return &NestedProvider{provider: provider}
}

func (fp *NestedProvider) Provide() (Quote, error) {
	q, err := fp.provider.Provide()
	if err != nil {
		return Quote{}, fmt.Errorf("file provider failed due to %w", err)
	}

	return q, nil
}
