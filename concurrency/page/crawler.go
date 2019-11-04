package page

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

type Image struct {
	Path string
}
type Images []Image

type Crawler interface {
	Crawl(url string) (Images, error)
}

type WebsiteCrawler struct {
	logger *logrus.Logger
}

func (wc *WebsiteCrawler) Crawl(url string) (Images, error) {
	var i Images
	wc.logger.Println("crawling...", url)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	// go to url
	// find property link
	// go to property details
	// find images
	// download image
	// return images and their local paths
	return i, nil
}

func NewWebsiteCrawler(logger *logrus.Logger) *WebsiteCrawler {
	return &WebsiteCrawler{logger: logger}
}
