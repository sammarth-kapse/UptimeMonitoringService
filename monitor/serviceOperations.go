package monitor

import (
	"UptimeMonitoringService/database"
)

func GetURLDataByID(id string) (database.URLData, bool) {
	return database.GetURLDataByID(id)
}

// Updates frequency/ status for corresponding request.
func UpdateURL(id string, request URLPatchRequest) (database.URLData, bool) {

	urlInfo, isPresent := GetURLDataByID(id)
	if !isPresent {
		return database.URLData{}, false
	}

	if urlInfo.CheckIfURLStatusISActive() {
		stopMonitoring(&urlInfo)
	}

	urlInfo.UpdateFrequency(request.Frequency)
	urlInfo.UpdateStatus(request.Status)
	urlInfo.ResetFailureCount()

	if urlInfo.Status == ACTIVE {
		go monitor(&urlInfo)
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

	urlInfo.UpdateStatus(ACTIVE)
	urlInfo.ResetFailureCount()
	go monitor(&urlInfo)

	return urlInfo.URL, true, false
}

func DeactivateURL(id string) (string, bool, bool) {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == INACTIVE {
		return urlInfo.URL, true, true
	}

	stopMonitoring(&urlInfo)

	return urlInfo.URL, true, false
}

func DeleteURLData(id string) bool {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return false
	}

	if urlInfo.CheckIfURLStatusISActive() {
		stopMonitoring(&urlInfo)
	}

	urlInfo.RemoveURLDataFromDatabase()
	return true
}
