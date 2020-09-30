package monitor

import (
	"UptimeMonitoringService/database"
	"fmt"
	"regexp"
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
	Frequency        int `json:"frequency"`
	FailureThreshold int `json:"failure_threshold"`
	CrawlTimeout     int `json:"crawl_timeout"`
}

func isURLStatusActive(urlInfo *database.URLData) bool {

	err := repository.DatabaseGet(urlInfo)
	handleError(err)
	return urlInfo.Status == ACTIVE
}

func increaseFailureCount(urlInfo *database.URLData) {
	urlInfo.FailureCount++
	if urlInfo.FailureCount >= urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	err := repository.DatabaseSave(urlInfo)
	handleError(err)
}

func formatURLProtocol(urlAddress string) string {

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

func checkIfURLEmpty(urlInfo database.URLData) bool {
	return urlInfo.URL == ""
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
