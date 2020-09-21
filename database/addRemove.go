package database

func (urlInfo *UrlData) AddURLDataInDatabase() {
	db.Create(&urlInfo)
}

func (urlInfo *UrlData) RemoveURLDataFromDatabase() {
	db.Delete(&urlInfo)
}
