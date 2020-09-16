package Monitor

type UrlPostRequest struct {
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	FailureThreshold int    `json:"failure_threshold"`
}

type UrlPatchRequest struct {
	Frequency int    `json:"frequency"`
	Status    string `json:"status"`
}

type UrlStatus struct {
	ID               string
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}
