package monitor

import "fmt"

const ACTIVE = "active"
const INACTIVE = "inactive"

type URLPostRequest struct {
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	FailureThreshold int    `json:"failure_threshold"`
}

type URLPatchRequest struct {
	Frequency        int `json:"frequency"`
	FailureThreshold int `json:"failure_threshold"`
	CrawlTimeout     int `json:"crawl_timeout"`
}

type URLData struct {
	ID               string `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

func checkIfURLEmpty(urlInfo URLData) bool {
	return urlInfo.URL == ""
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
