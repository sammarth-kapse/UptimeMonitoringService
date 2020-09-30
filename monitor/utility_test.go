package monitor

import (
	"UptimeMonitoringService/database"
	"UptimeMonitoringService/mocks"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestIncreaseFailureCount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseSave(gomock.Any()).Return(nil).MaxTimes(2)
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	urlInfo := database.URLData{
		FailureThreshold: 5,
		Status:           ACTIVE,
		FailureCount:     0,
	}

	// when failure_count + 1 != failure_threshold:
	increaseFailureCount(&urlInfo)
	assert.Equal(t, urlInfo.FailureCount, 1)
	assert.Equal(t, urlInfo.Status, ACTIVE)

	// when failure_count + 1 == failure_threshold:
	urlInfo.FailureCount = 4
	increaseFailureCount(&urlInfo)
	assert.Equal(t, urlInfo.FailureCount, 5)
	assert.Equal(t, urlInfo.Status, INACTIVE)
}

func TestIsURLStatusActiveWhenActive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(&database.URLData{}).DoAndReturn(func(urlInfo *database.URLData) error {
		urlInfo.Status = ACTIVE
		return nil
	})
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	isActive := isURLStatusActive(&database.URLData{})
	assert.Equal(t, isActive, true)
}

func TestIsURLStatusActiveWhenInactive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)

	mockRepo.EXPECT().DatabaseGet(&database.URLData{}).DoAndReturn(func(urlInfo *database.URLData) error {
		urlInfo.Status = INACTIVE
		return nil
	})
	database.SetRepoController(mockRepo)
	repository = database.GetRepoController()

	isActive := isURLStatusActive(&database.URLData{})
	assert.Equal(t, isActive, false)
}

func TestFormatURLProtocol(t *testing.T) {
	type testFormat struct {
		url            string
		expectedResult string
	}

	testArr := []testFormat{
		{"https://testURL.com", "https://testURL.com"},
		{"https://www.testURL.com", "https://www.testURL.com"},
		{"http://testURL.com", "http://testURL.com"},
		{"http://www.testURL.com", "http://www.testURL.com"},
		{"testURL.com", "https://testURL.com"},
		{"www.testURL.com", "https://www.testURL.com"},
	}

	for _, testItem := range testArr {
		actualResult := formatURLProtocol(testItem.url)
		if testItem.expectedResult != actualResult {
			t.Error("For urlInfo: ", testItem.url, " Expected Result: ", testItem.expectedResult, " Got: ",
				actualResult)
		}
	}
}

func TestIsHttpOrHttpsRequest(t *testing.T) {
	type testFormat struct {
		url            string
		expectedResult bool
	}

	testArr := []testFormat{
		{"https://testURL.com", true},
		{"https://www.testURL.com", true},
		{"http://testURL.com", true},
		{"http://www.testURL.com", true},
		{"testURL.com", false},
		{"www.testURL.com", false},
	}

	for _, testItem := range testArr {
		actualResult := isHttpOrHttpsRequest(testItem.url)
		if testItem.expectedResult != actualResult {
			t.Error("For urlInfo: ", testItem.url, " Expected Result: ", testItem.expectedResult, " Got: ",
				actualResult)
		}
	}
}

func TestCheckIfURLEmpty(t *testing.T) {

	type testFormat struct {
		urlInfo        database.URLData
		expectedResult bool
	}

	testArr := []testFormat{
		{database.URLData{URL: "http://testURL.com"}, false},
		{database.URLData{}, true}, // Empty URL
	}

	for _, testItem := range testArr {
		actualResult := checkIfURLEmpty(testItem.urlInfo)
		if testItem.expectedResult != actualResult {
			t.Error("For urlInfo: ", testItem.urlInfo, " Expected Result: ", testItem.expectedResult, " Got: ",
				actualResult)
		}
	}
}
