package hello

import (
	"html/template"
	"net/http"
	"sync"
)

type Handler2 struct{}

func NewHandler2() Handler2 {
	return Handler2{}
}

func (h Handler2) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var (
		init      sync.Once
		templates *template.Template
		err       error
	)
	init.Do(func() {
		templates = template.Must(template.ParseGlob("web/templates/*"))
	})

	provider := NewFileProvider("./file")
	quotes, err := provider.Provide()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	viewModel := struct {
		Quotes Quotes
	}{
		Quotes: quotes,
	}

	if err = templates.ExecuteTemplate(res, "index.html", viewModel); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
