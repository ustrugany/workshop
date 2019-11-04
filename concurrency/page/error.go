package page

type CrawlerError struct {
	Reason string
	Url    string
	Err    error
}

func (ce CrawlerError) Error() string {
	return ce.Reason
}
