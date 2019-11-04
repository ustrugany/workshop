package page

import (
	"errors"
	"log"
	"net/http"
	url "net/url"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type CrawlHandler struct {
	logger *logrus.Logger
}

func (ch *CrawlHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// parse uri
	vars := mux.Vars(req)
	urlToCrawl, ok := vars["url"]
	if !ok {
		http.Error(res, "missing url query param", http.StatusNotFound)
		return
	}

	u, err := url.Parse(urlToCrawl)
	if err != nil {
		ch.logger.Fatalf("could not parse url to crawl %q", err)
		http.Error(res, "invalid url to parse", http.StatusNotFound)
		return
	}

	numberOfPages := 5
	for i := 1; i <= numberOfPages; i++ {
		q := u.Query()
		q.Set("page", strconv.Itoa(i))
		u.RawQuery = q.Encode()
		ch.logger.Println(u.String())

		c := NewWebsiteCrawler(ch.logger)
		i, err := c.Crawl(u.String())
		var e CrawlerError
		if err != nil {
			if errors.As(err, &e) {
				ch.logger.Errorf("could not crawl url %s due to %s", e.Url, e.Reason)
			} else {
				ch.logger.Fatalf("unexpected error [%v]", err)
			}
		}

		ch.logger.Println(i)
	}

	if _, err := res.Write([]byte("")); err != nil {
		log.Fatal(err)
	}
}

func NewCrawlHandler(logger *logrus.Logger) *CrawlHandler {
	return &CrawlHandler{logger: logger}
}
