package monitor

import "github.com/google/uuid"

func AddService(req URLPostRequest) (URLData, error) {

	req.URL = formatURLProtocol(req.URL)

	id := uuid.New().String()

	urlInfo := URLData{
		ID:               id,
		URL:              req.URL,
		CrawlTimeout:     req.CrawlTimeout,
		Frequency:        req.Frequency,
		FailureThreshold: req.FailureThreshold,
		Status:           ACTIVE,
		FailureCount:     0,
	}

	err := repository.databaseCreate(&urlInfo)
	handleError(err)

	go monitor(&urlInfo)

	return urlInfo, nil
}

func GetURLDataByID(id string) (*URLData, bool) {

	urlInfo := URLData{
		ID: id,
	}
	err := repository.databaseGet(&urlInfo)
	handleError(err)
	if checkIfURLEmpty(urlInfo) { // when url == empty -> id is invalid; In utility.go
		return &URLData{}, false
	}
	return &urlInfo, true
}

func UpdateURL(id string, request URLPatchRequest) (URLData, bool) {

	urlInfo, isPresent := GetURLDataByID(id)
	if !isPresent {
		return URLData{}, false
	}

	if isURLStatusActive(urlInfo) {
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
	err := repository.databaseSave(urlInfo)
	handleError(err)

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
	err := repository.databaseSave(urlInfo)
	handleError(err)

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

	if isURLStatusActive(urlInfo) {
		stopMonitoring(urlInfo)
	}
	err := repository.databaseDelete(urlInfo)
	handleError(err)

	return true
}
