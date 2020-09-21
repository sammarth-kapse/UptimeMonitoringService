package database

func GetURLDataByID(id string) (UrlData, bool) {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	if checkIfURLEmpty(urlInfo) { // when url == empty id is invalid
		return urlInfo, false
	}
	return urlInfo, true
}

func CheckIfURLStatusISActive(id string) bool {

	var urlInfo UrlData
	urlInfo.getURLInfoFromDatabase(id)
	return urlInfo.Status == ACTIVE
}
