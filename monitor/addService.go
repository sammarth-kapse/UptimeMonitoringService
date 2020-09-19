package monitor

import (
	"UptimeMonitoringService/database"
	"github.com/google/uuid"
	"regexp"
)

// 'url' is used as package ('net/url') not a variable. urlAddress is the variable used instead.
// 'URL' is a field of struct URLData

func AddService(req URLPostRequest) (database.UrlData, error) {

	req.URL = checkForProtocolInURL(req.URL)

	id := uuid.New().String()

	newURLData := database.UrlData{
		ID:               id,
		URL:              req.URL,
		CrawlTimeout:     req.CrawlTimeout,
		Frequency:        req.Frequency,
		FailureThreshold: req.FailureThreshold,
		Status:           ACTIVE,
		FailureCount:     0,
	}

	database.AddURLDataInDatabase(newURLData)

	go monitor(newURLData.ID, newURLData.URL, newURLData.Frequency, newURLData.CrawlTimeout)

	return newURLData, nil
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
