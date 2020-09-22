package database

// The only function in database pkg which is not a method of type URLData
// This is because the function returns url data for the corresponding id.
func GetURLDataByID(id string) (URLData, bool) {

	//var urlInfo URLData
	urlInfo := URLData{
		ID: id,
	}
	urlInfo.getURLInfoFromDatabase()
	if checkIfURLEmpty(urlInfo) { // when url == empty -> id is invalid
		return urlInfo, false
	}
	return urlInfo, true
}

func (urlInfo *URLData) CheckIfURLStatusISActive() bool {

	urlInfo.getURLInfoFromDatabase()
	return urlInfo.Status == ACTIVE
}
