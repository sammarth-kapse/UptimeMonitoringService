package database

func IncreaseFailureCount(id string) {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)

	urlInfo.FailureCount++
	if urlInfo.FailureCount == urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	urlInfo.saveIntoDatabase()
}

func SetUrlAsInactive(id string) {
	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	urlInfo.Status = INACTIVE
	urlInfo.saveIntoDatabase()
}

func UpdateFrequency(id string, newFrequency int) {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	urlInfo.Frequency = newFrequency
	urlInfo.saveIntoDatabase()
}

func UpdateStatus(id string, newStatus string) {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	urlInfo.Status = newStatus
	urlInfo.saveIntoDatabase()
}

func ResetFailureCount(id string) {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	urlInfo.FailureCount = 0
	urlInfo.saveIntoDatabase()
}
