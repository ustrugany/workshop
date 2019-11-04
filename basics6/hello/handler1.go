package hello

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler1 struct{}

func NewHandler1() Handler1 {
	return Handler1{}
}

func (h Handler1) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	provider := NewFileProvider("./file")
	q, err := provider.Provide()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	} else {
		body, _ := json.Marshal(q)
		if _, err := res.Write(body); err != nil {
			log.Fatal(err)
		}
	}
}
