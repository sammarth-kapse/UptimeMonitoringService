package monitor

import "UptimeMonitoringService/database"

type RepositoryInterface interface {
	DatabaseGet(urlInfo *URLData)
	DatabaseSave(urlInfo *URLData)
	DatabaseCreate(urlInfo *URLData)
	DatabaseDelete(urlInfo *URLData)
}

type RepoStruct struct{}

var Repository RepositoryInterface

func SetRepo(repoType RepositoryInterface) {
	Repository = repoType
}

func (rp *RepoStruct) DatabaseGet(urlInfo *URLData) {
	database.DB.First(&urlInfo)
}

func (rp *RepoStruct) DatabaseSave(urlInfo *URLData) {
	database.DB.Save(&urlInfo)
}

func (rp *RepoStruct) DatabaseCreate(urlInfo *URLData) {
	database.DB.Create(&urlInfo)
}

func (rp *RepoStruct) DatabaseDelete(urlInfo *URLData) {
	database.DB.Delete(&urlInfo)
}
