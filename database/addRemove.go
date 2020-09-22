package database

func (urlInfo *URLData) AddURLDataInDatabase() {
	db.Create(&urlInfo)
}

func (urlInfo *URLData) RemoveURLDataFromDatabase() {
	db.Delete(&urlInfo)
}
