package monitor

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"regexp"
	"time"
)

// 'url' is used as package ('net/url') not a variable. urlAddress is the variable used instead.
// 'URL' is a field of struct URLData

func AddService(req URLPostRequest) (*URLData, error) {

	req.URL = checkForProtocolInURL(req.URL)

	id := uuid.New().String()

	newURLData := new(URLData)
	newURLData.ID = id
	newURLData.URL = req.URL
	newURLData.CrawlTimeout = req.CrawlTimeout
	newURLData.Frequency = req.Frequency
	newURLData.FailureThreshold = req.FailureThreshold
	newURLData.Status = ACTIVE
	newURLData.FailureCount = 0

	insertIntoURLCollection(id, newURLData)

	go monitor(newURLData)

	return newURLData, nil
}

// urlInfo would be used as a variable of type UrlData
func monitor(urlInfo *URLData) {

	for urlInfo.FailureCount < urlInfo.FailureThreshold && urlInfo.Status == ACTIVE {

		resp, err := http.Get(urlInfo.URL)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			urlInfo.FailureCount++
		}

		fmt.Println("Checked Url : ", urlInfo.URL, " is: ", urlInfo.Status, " code: ", resp.StatusCode)
		if urlInfo.FailureCount < urlInfo.FailureThreshold && urlInfo.Status == ACTIVE {
			time.Sleep(time.Duration(urlInfo.Frequency) * time.Second)
		}
	}
	urlInfo.Status = INACTIVE
	fmt.Println("Url: ", urlInfo.URL, " is now: ", urlInfo.Status)
}

func checkForProtocolInURL(urlAddress string) string {

	if !isHttpOrHttpsRequest(urlAddress) {
		urlAddress = "https://" + urlAddress
	}
	return urlAddress
}

func isHttpOrHttpsRequest(urlAddress string) bool {
	isHttp, err := regexp.MatchString("http://([a-z]+)", urlAddress)
	handleError(err)
	var isHttps bool
	isHttps, err = regexp.MatchString("https://([a-z]+)", urlAddress)
	return isHttp || isHttps
}
