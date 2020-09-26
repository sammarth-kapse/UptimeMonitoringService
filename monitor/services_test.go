package monitor

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetURLDataByIDWhenValidID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().databaseGet(&URLData{ID: "testID"}).DoAndReturn(func(urlInfo *URLData) error {
		urlInfo.URL = "https://testURL.com"
		return nil
	})
	setRepoController(mockRepo)

	urlInfo, isPresent := GetURLDataByID("testID")
	assert.Equal(t, urlInfo, &URLData{ID: "testID", URL: "https://testURL.com"})
	assert.Equal(t, isPresent, true)
}

func TestGetURLDataByIDWhenInvalidID(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().databaseGet(gomock.Any()).Return(nil)
	setRepoController(mockRepo)

	urlInfo, isPresent := GetURLDataByID("testID")
	assert.Equal(t, urlInfo, &URLData{})
	assert.Equal(t, isPresent, false)
}
