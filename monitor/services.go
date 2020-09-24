package monitor

func UpdateURL(id string, request URLPatchRequest) (URLData, bool) {

	urlInfo, isPresent := GetURLDataByID(id)
	if !isPresent {
		return URLData{}, false
	}

	if CheckIfURLStatusISActive(urlInfo) {
		stopMonitoring(urlInfo)
	}

	if request.Frequency != 0 {
		urlInfo.Frequency = request.Frequency
	}
	if request.FailureThreshold != 0 {
		urlInfo.FailureThreshold = request.FailureThreshold
	}
	if request.CrawlTimeout != 0 {
		urlInfo.CrawlTimeout = request.CrawlTimeout
	}
	urlInfo.Status = ACTIVE
	urlInfo.FailureCount = 0
	Repository.DatabaseSave(urlInfo)

	go monitor(urlInfo)

	return *urlInfo, true
}

func ActivateURL(id string) (string, bool, bool) {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == ACTIVE {
		return urlInfo.URL, true, true
	}

	urlInfo.Status = ACTIVE
	urlInfo.FailureCount = 0
	Repository.DatabaseSave(urlInfo)

	go monitor(urlInfo)

	return urlInfo.URL, true, false
}

func DeactivateURL(id string) (string, bool, bool) {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return "", false, false
	} else if urlInfo.Status == INACTIVE {
		return urlInfo.URL, true, true
	}

	stopMonitoring(urlInfo)

	return urlInfo.URL, true, false
}

func DeleteURLData(id string) bool {

	urlInfo, isPresent := GetURLDataByID(id)

	if !isPresent {
		return false
	}

	if CheckIfURLStatusISActive(urlInfo) {
		stopMonitoring(urlInfo)
	}
	Repository.DatabaseDelete(urlInfo)

	return true
}
