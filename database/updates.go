package database

func IncreaseFailureCount(id string) {

	urlInfo := getURLInfoInternal(id)

	urlInfo.FailureCount++
	if urlInfo.FailureCount == urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	db.Save(&urlInfo)
}

func SetUrlAsInactive(id string) {
	urlInfo := getURLInfoInternal(id)
	urlInfo.Status = INACTIVE
	db.Save(&urlInfo)
}

func UpdateFrequency(id string, newFrequency int) {

	urlInfo := getURLInfoInternal(id)
	urlInfo.Frequency = newFrequency
	db.Save(&urlInfo)
}

func UpdateStatus(id string, newStatus string) {

	urlInfo := getURLInfoInternal(id)
	urlInfo.Status = newStatus
	db.Save(&urlInfo)
}

func ResetFailureCount(id string) {

	urlInfo := getURLInfoInternal(id)
	urlInfo.FailureCount = 0
	db.Save(&urlInfo)
}
