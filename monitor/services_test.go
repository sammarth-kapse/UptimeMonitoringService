package monitor

import (
	"UptimeMonitoringService/database"
	"UptimeMonitoringService/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetURLDataByIDWhenValidID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(&database.URLData{ID: "testID"}).DoAndReturn(func(urlInfo *database.URLData) error {
		urlInfo.URL = "https://testURL.com"
		return nil
	})
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	urlInfo, isPresent := GetURLDataByID("testID")
	assert.Equal(t, urlInfo, &database.URLData{ID: "testID", URL: "https://testURL.com"})
	assert.Equal(t, isPresent, true)
}

func TestGetURLDataByIDWhenInvalidID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(gomock.Any()).Return(nil)
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	urlInfo, isPresent := GetURLDataByID("testID")
	assert.Equal(t, urlInfo, &database.URLData{})
	assert.Equal(t, isPresent, false)
}
