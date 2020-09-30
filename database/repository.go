package database

import (
	"errors"
)

type RepositoryController interface {
	DatabaseGet(urlInfo *URLData) error
	DatabaseSave(urlInfo *URLData) error
	DatabaseCreate(urlInfo *URLData) error
	DatabaseDelete(urlInfo *URLData) error
}

type MonitorRepo struct{}

var repository RepositoryController

func SetRepoController(repoType RepositoryController) {
	repository = repoType
}

func GetRepoController() RepositoryController {
	return repository
}

func (rp *MonitorRepo) DatabaseCreate(urlInfo *URLData) error {
	return DB.Create(&urlInfo).Error
}

func (rp *MonitorRepo) DatabaseGet(urlInfo *URLData) error {
	result := DB.First(&urlInfo)
	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}
	return result.Error
}

func (rp *MonitorRepo) DatabaseSave(urlInfo *URLData) error {
	return DB.Save(&urlInfo).Error
}

func (rp *MonitorRepo) DatabaseDelete(urlInfo *URLData) error {
	result := DB.Delete(&urlInfo)
	if result.RowsAffected == 0 {
		return errors.New("no record found")
	}
	return result.Error
}
