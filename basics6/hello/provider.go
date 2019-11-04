package hello

import (
	"bufio"
	"fmt"
	"os"
)

type Quote struct {
	Content string `json:"quote"`
}

type Quotes []Quote

type Provider interface {
	Provide() (Quotes, error)
}

type FileProvider struct {
	path string
}

func NewFileProvider(path string) *FileProvider {
	return &FileProvider{path: path}
}

func (fp *FileProvider) Provide() (Quotes, error) {
	var quotes []Quote
	f, err := os.Open(fp.path)
	if err != nil {
		return quotes, ProviderError{Reason: fmt.Sprintf("file does not exists %q", fp.path), Err: err}
	}

	defer func() {
		_ = f.Close()
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		quotes = append(quotes, Quote{Content: scanner.Text()})
	}
	if err := scanner.Err(); err != nil {
		return quotes, ProviderError{Reason: fmt.Sprintf("failed to scan file %q", fp.path), Err: err}
	}

	return quotes, nil
}
