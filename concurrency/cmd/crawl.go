package cmd

import (
	"errors"
	"net/url"
	"strconv"
	"sync"

	"example.com/concurrency/page"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newCrawlerHandler() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		logger := logrus.New()
		logger.Formatter = new(logrus.JSONFormatter)
		urlFlag, _ := cmd.Flags().GetString("url")
		concurrencyFlag, _ := cmd.Flags().GetInt("concurrency")
		u, err := url.Parse(urlFlag)
		if err != nil {
			logger.Fatalf("could not parse url to crawl %q", err)
		}

		wg := &sync.WaitGroup{}
		numberOfPages := concurrencyFlag
		for index := 1; index <= numberOfPages; index++ {
			q := u.Query()
			q.Set("page", strconv.Itoa(index))
			u.RawQuery = q.Encode()
			logrus.Println(u.String())

			wg.Add(1)
			go func(wg *sync.WaitGroup, index int) {
				c := page.NewWebsiteCrawler(logger)
				img, err := c.Crawl(u.String())
				var e page.CrawlerError
				if err != nil {
					if errors.As(err, &e) {
						logger.Errorf("could not crawl url %s due to %s", e.Url, e.Reason)
					} else {
						logger.Fatalf("unexpected error [%v]", err)
					}
				}
				logger.Println("crawler", index, "finished")
				logger.Println(img)
				wg.Done()
			}(wg, index)
		}
		wg.Wait()
	}
}

func NewCrawlerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "crawler",
		Short: "crawler command",
		Run:   newCrawlerHandler(),
	}
}
