package monitor

func GetURLDataByID(id string) (*URLData, bool) {

	urlInfo := URLData{
		ID: id,
	}
	Repository.DatabaseGet(&urlInfo)
	if checkIfURLEmpty(urlInfo) { // when url == empty -> id is invalid; In utility.go
		return &urlInfo, false
	}
	return &urlInfo, true
}

func CheckIfURLStatusISActive(urlInfo *URLData) bool {

	Repository.DatabaseGet(urlInfo)
	return urlInfo.Status == ACTIVE
}

func IncreaseFailureCount(urlInfo *URLData) {
	urlInfo.FailureCount++
	if urlInfo.FailureCount >= urlInfo.FailureThreshold {
		urlInfo.Status = INACTIVE
	}
	Repository.DatabaseSave(urlInfo)
}
