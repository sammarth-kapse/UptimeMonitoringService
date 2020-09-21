package database

func GetURLDataByID(id string) (UrlData, bool) {

	//var urlInfo UrlData
	urlInfo := UrlData{
		ID: id,
	}
	db.First(&urlInfo)
	if checkIfURLEmpty(urlInfo) { // when url == empty id is invalid
		return urlInfo, false
	}
	return urlInfo, true
}

func (urlInfo *UrlData) CheckIfURLStatusISActive() bool {

	urlInfo.getURLInfoFromDatabase()
	return urlInfo.Status == ACTIVE
}
