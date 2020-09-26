package monitor

import (
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

type URLData struct {
	ID               string `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

func isURLStatusActive(urlInfo *URLData) bool {

	err := repository.databaseGet(urlInfo)
	handleError(err)
	return urlInfo.Status == ACTIVE
}

func increaseFailureCount(urlInfo *URLData) {
	urlInfo.FailureCount++
	if urlInfo.FailureCount >= urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	err := repository.databaseSave(urlInfo)
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

func checkIfURLEmpty(urlInfo URLData) bool {
	return urlInfo.URL == ""
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
