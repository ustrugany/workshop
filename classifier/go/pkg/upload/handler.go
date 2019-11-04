package upload

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type UploadHandler struct {
	Path string
}

func (h *UploadHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		t, _ := template.ParseFiles("web/upload.gtpl")
		if err := t.Execute(res, nil); err != nil {
			log.Fatal(err)
		}
	} else if req.Method == "POST" {
		// Post
		file, handler, err := req.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Println(err)
			}
		}()

		path := fmt.Sprintf("%s/%s", h.Path, handler.Filename)
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
			}
		}()

		if _, err := io.Copy(f, file); err != nil {
			log.Fatal(err)
		}

		result := h.Classifier.Classify(path, 10*time.Second)
		if _, err := res.Write([]byte(result)); err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Unknown HTTP " + req.Method + "  Method")
	}
}
