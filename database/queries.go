package database

func GetURLDataByID(id string) (UrlData, bool) {

	urlInfo := *getURLInfoInternal(id)
	if checkIfURLEmpty(urlInfo) { // when url == empty id is invalid
		return urlInfo, false
	}
	return urlInfo, true
}

func CheckIfURLStatusISActive(id string) bool {

	urlInfo := getURLInfoInternal(id)
	return urlInfo.Status == ACTIVE
}

func getURLInfoInternal(id string) *UrlData {
	urlInfo := UrlData{
		ID: id,
	}
	db.First(&urlInfo)
	return &urlInfo
}
