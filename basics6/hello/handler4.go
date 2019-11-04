package hello

import (
	"html/template"
	"net/http"
	"sync"
)

type Handler4 struct{}

func NewHandler4() Handler4 {
	return Handler4{}
}

func (h Handler4) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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

	if err = templates.ExecuteTemplate(res, "index2.html", viewModel); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
