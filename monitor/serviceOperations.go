package monitor

import (
	"UptimeMonitoringService/database"
)

func GetURLDataByID(id string) (database.UrlData, bool) {
	return database.GetURLDataByID(id)
}

func UpdateURL(id string, request URLPatchRequest) (database.UrlData, bool) {

	urlInfo, isPresent := GetURLDataByID(id)
	if !isPresent {
		return database.UrlData{}, false
	}

	stopMonitoring(id)

	database.UpdateFrequency(id, request.Frequency)
	database.UpdateStatus(id, request.Status)
	database.ResetFailureCount(id)

	urlInfo, _ = GetURLDataByID(id)
	if urlInfo.Status == ACTIVE {
		go monitor(urlInfo.ID, urlInfo.URL, urlInfo.Frequency, urlInfo.CrawlTimeout)
	}
	return urlInfo, true
}

func ActivateURL(id string) (string, bool, bool) {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == ACTIVE {
		return urlInfo.URL, true, true
	}

	database.UpdateStatus(id, ACTIVE)
	database.ResetFailureCount(id)
	go monitor(urlInfo.ID, urlInfo.URL, urlInfo.Frequency, urlInfo.CrawlTimeout)

	return urlInfo.URL, true, false
}

func DeactivateURL(id string) (string, bool, bool) {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == INACTIVE {
		return urlInfo.URL, true, true
	}

	stopMonitoring(id)

	return urlInfo.URL, true, false
}

func DeleteURLData(id string) bool {

	_, isPresent := GetURLDataByID(id)

	if !isPresent {
		return false
	}

	stopMonitoring(id)
	database.RemoveURLDataFromDatabase(id)

	return true
}
