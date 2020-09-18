package monitor

import (
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"net/http"
	"time"
)

const ACTIVE = "active"
const INACTIVE = "inactive"

type URLPostRequest struct {
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	FailureThreshold int    `json:"failure_threshold"`
}

type URLPatchRequest struct {
	Frequency int    `json:"frequency"`
	Status    string `json:"status"`
}

type URLData struct {
	ID               string
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

// urlInfo would be used as a variable of type UrlData
func monitor(urlInfo *URLData) {

	isFirstCheck := true

	for urlInfo.FailureCount < urlInfo.FailureThreshold && urlInfo.Status == ACTIVE {

		if !isFirstCheck {
			ticker := time.NewTimer(time.Duration(urlInfo.Frequency) * time.Second)
			<-ticker.C
		}
		fmt.Println("###")
		go checkURL(urlInfo)
	}
	urlInfo.Status = INACTIVE
	fmt.Println("Url: ", urlInfo.URL, " is now: ", urlInfo.Status)
}

func stopMonitoring(urlInfo *URLData) {
	urlInfo.Status = INACTIVE
	time.Sleep(time.Duration(urlInfo.Frequency) * time.Second)
}

func checkURL(urlInfo *URLData) {

	if urlInfo.Status == INACTIVE {
		return
	}

	timeout := time.Duration(urlInfo.CrawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	resp, err := client.Get(urlInfo.URL, nil)
	if err != nil {
		urlInfo.FailureCount++
		return
	}
	if resp.StatusCode != http.StatusOK {
		urlInfo.FailureCount++
	}
	fmt.Println("Checked Url : ", urlInfo.URL, " is: ", urlInfo.Status, " code: ", resp.StatusCode)
}
