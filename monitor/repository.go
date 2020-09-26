package monitor

import "UptimeMonitoringService/database"

type RepositoryController interface {
	databaseGet(urlInfo *URLData) error
	databaseSave(urlInfo *URLData) error
	databaseCreate(urlInfo *URLData) error
	databaseDelete(urlInfo *URLData) error
}

type monitorRepo struct{}

var repository RepositoryController

func setRepoController(repoType RepositoryController) {
	repository = repoType
}

func (rp *monitorRepo) databaseGet(urlInfo *URLData) error {
	return database.DB.First(&urlInfo).Error
}

func (rp *monitorRepo) databaseSave(urlInfo *URLData) error {
	return database.DB.Save(&urlInfo).Error
}

func (rp *monitorRepo) databaseCreate(urlInfo *URLData) error {
	return database.DB.Create(&urlInfo).Error
}

func (rp *monitorRepo) databaseDelete(urlInfo *URLData) error {
	return database.DB.Delete(&urlInfo).Error
}
