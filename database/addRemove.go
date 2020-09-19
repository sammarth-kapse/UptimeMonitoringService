package database

func AddURLDataInDatabase(urlInfo UrlData) {
	db.Create(&urlInfo)
}

func RemoveURLDataFromDatabase(id string) {
	urlInfo := UrlData{
		ID: id,
	}
	db.Delete(&urlInfo)
}
