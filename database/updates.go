package database

func (urlInfo *URLData) IncreaseFailureCount() {

	urlInfo.getURLInfoFromDatabase()

	urlInfo.FailureCount++
	if urlInfo.FailureCount == urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	urlInfo.saveIntoDatabase()
}

func (urlInfo *URLData) SetUrlAsInactive() {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Status = INACTIVE
	urlInfo.saveIntoDatabase()
}

func (urlInfo *URLData) UpdateFrequency(newFrequency int) {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Frequency = newFrequency
	urlInfo.saveIntoDatabase()
}

func (urlInfo *URLData) UpdateStatus(newStatus string) {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.Status = newStatus
	urlInfo.saveIntoDatabase()
}

func (urlInfo *URLData) ResetFailureCount() {

	urlInfo.getURLInfoFromDatabase()
	urlInfo.FailureCount = 0
	urlInfo.saveIntoDatabase()
}
