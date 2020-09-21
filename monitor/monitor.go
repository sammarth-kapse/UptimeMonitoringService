package monitor

import (
	"UptimeMonitoringService/database"
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

// Monitors a URL till it's status is 'active'
func monitor(urlInfo *database.UrlData) {

	isFirstCheck := true

	for urlInfo.CheckIfURLStatusISActive() {
		if !isFirstCheck {
			ticker := time.NewTimer(time.Duration(urlInfo.Frequency) * time.Second)
			<-ticker.C
		}
		go checkURL(urlInfo)
		isFirstCheck = false
	}
}

func stopMonitoring(urlInfo *database.UrlData) {

	urlInfo.SetUrlAsInactive()
	time.Sleep(10 * time.Second)
}

// Makes an HTTP GET request on the url with the given set of parameters.
func checkURL(urlInfo *database.UrlData) {

	if !urlInfo.CheckIfURLStatusISActive() {
		return
	}

	timeout := time.Duration(urlInfo.CrawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	resp, err := client.Get(urlInfo.URL, nil)
	if err != nil {
		urlInfo.IncreaseFailureCount() // Request didn't complete within crawl_timeout.
		return
	}

	if resp.StatusCode != http.StatusOK {
		urlInfo.IncreaseFailureCount() // Unexpected status-code in response.
	}

	fmt.Println("Checked Url : ", urlInfo.URL, " code: ", resp.StatusCode, "   time: ", time.Now())
}
