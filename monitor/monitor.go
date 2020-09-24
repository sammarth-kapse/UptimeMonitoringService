package monitor

import (
	"UptimeMonitoringService/database"
	"fmt"
	"github.com/gojektech/heimdall/httpclient"
	"net/http"
	"os"
	"time"
)

// All the structures and constants are present in utility.go

func init() {
	err := database.DB.AutoMigrate(&URLData{})
	if err != nil {
		os.Exit(1)
	}
	SetRepo(&RepoStruct{})
}

// Monitors a URL till it's status is 'active'
func monitor(urlInfo *URLData) {

	isFirstCheck := true

	for CheckIfURLStatusISActive(urlInfo) {
		if !isFirstCheck {
			ticker := time.NewTimer(time.Duration(urlInfo.Frequency) * time.Second)
			<-ticker.C
		}
		go checkURL(urlInfo)
		isFirstCheck = false
	}
}

func stopMonitoring(urlInfo *URLData) {

	urlInfo.Status = INACTIVE
	Repository.DatabaseSave(urlInfo)

	time.Sleep(time.Duration(urlInfo.Frequency+2) * time.Second)
}

// Makes an HTTP GET request on the url with the given set of parameters.
func checkURL(urlInfo *URLData) {

	if !CheckIfURLStatusISActive(urlInfo) {
		return
	}

	resp, err := makeHTTPGetRequest(urlInfo.CrawlTimeout, urlInfo.URL)

	if err != nil {
		IncreaseFailureCount(urlInfo) // Request didn't complete within crawl_timeout.
		return
	}

	if resp.StatusCode != http.StatusOK {
		IncreaseFailureCount(urlInfo) // Unexpected status-code in response.
	}

	fmt.Println("Checked Url : ", urlInfo.URL, " code: ", resp.StatusCode, "   time: ", time.Now())
}

func makeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error) {

	timeout := time.Duration(crawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	return client.Get(url, nil)
}
