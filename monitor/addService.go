package monitor

import (
	"UptimeMonitoringService/database"
	"github.com/google/uuid"
	"regexp"
)

// Adds the URL data into the database and starts monitoring.
func AddService(req URLPostRequest) (database.UrlData, error) {

	req.URL = checkForProtocolInURL(req.URL)

	id := uuid.New().String()

	urlInfo := database.UrlData{
		ID:               id,
		URL:              req.URL,
		CrawlTimeout:     req.CrawlTimeout,
		Frequency:        req.Frequency,
		FailureThreshold: req.FailureThreshold,
		Status:           ACTIVE,
		FailureCount:     0,
	}

	urlInfo.AddURLDataInDatabase()

	go monitor(&urlInfo)

	return urlInfo, nil
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
