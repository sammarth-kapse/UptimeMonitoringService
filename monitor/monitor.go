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
func monitor(id, url string, frequency, crawlTimeout int) {

	isFirstCheck := true

	for database.CheckIfURLStatusISActive(id) {
		if !isFirstCheck {
			ticker := time.NewTimer(time.Duration(frequency) * time.Second)
			<-ticker.C
		}
		go checkURL(id, url, crawlTimeout)
		isFirstCheck = false
	}
}

func stopMonitoring(id string) {

	database.SetUrlAsInactive(id)
	time.Sleep(10 * time.Second)
}

// Makes an HTTP GET request on the url with the given set of parameters.
func checkURL(id, url string, crawlTimeout int) {

	if !database.CheckIfURLStatusISActive(id) {
		return
	}

	timeout := time.Duration(crawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	resp, err := client.Get(url, nil)
	if err != nil {
		database.IncreaseFailureCount(id) // Request didn't complete within crawl_timeout.
		return
	}

	if resp.StatusCode != http.StatusOK {
		database.IncreaseFailureCount(id) // Unexpected status-code in response.
	}

	fmt.Println("Checked Url : ", url, " code: ", resp.StatusCode, "   time: ", time.Now())
}
