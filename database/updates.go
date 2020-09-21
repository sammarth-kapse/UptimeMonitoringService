package database

func (urlInfo *UrlData) IncreaseFailureCount() {

	urlInfo.getURLInfoFromDatabase()

	urlInfo.FailureCount++
	if urlInfo.FailureCount == urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	urlInfo.saveIntoDatabase()
}

func (urlInfo *UrlData) SetUrlAsInactive() {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Status = INACTIVE
	urlInfo.saveIntoDatabase()
}

func (urlInfo *UrlData) UpdateFrequency(newFrequency int) {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Frequency = newFrequency
	urlInfo.saveIntoDatabase()
}

func (urlInfo *UrlData) UpdateStatus(newStatus string) {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Status = newStatus
	urlInfo.saveIntoDatabase()
}

func (urlInfo *UrlData) ResetFailureCount() {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.FailureCount = 0
	urlInfo.saveIntoDatabase()
}
