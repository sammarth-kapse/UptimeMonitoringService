package monitor

import (
	"UptimeMonitoringService/database"
	"UptimeMonitoringService/httpRequests"
	"UptimeMonitoringService/mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCheckURLUptime(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mocks.NewMockRepositoryController(ctrl)
	mockHttp := mocks.NewMockHttpController(ctrl)

	mockRepo.EXPECT().DatabaseSave(gomock.Any()).Return(nil).MaxTimes(2)
	mockRepo.EXPECT().DatabaseGet(gomock.Any()).Return(nil).MaxTimes(3)
	database.SetRepoController(mockRepo)

	mockHttp.EXPECT().MakeHTTPGetRequest(gomock.Any(), "https://case1.com").Return(
		&http.Response{StatusCode: http.StatusOK}, nil)

	mockHttp.EXPECT().MakeHTTPGetRequest(gomock.Any(), "https://case2.com").Return(
		&http.Response{StatusCode: http.StatusConflict}, nil)

	mockHttp.EXPECT().MakeHTTPGetRequest(gomock.Any(), "https://case3.com").Return(
		&http.Response{StatusCode: http.StatusConflict}, fmt.Errorf("SomeError"))

	httpRequests.SetHTTPController(mockHttp)

	repository = database.GetRepoController()
	httpCalls = httpRequests.GetHTTPController()

	// Case 1: status_code = 200 AND crawl_timeout > delay
	urlInfo1 := database.URLData{
		ID:               "Case1",
		FailureCount:     0,
		FailureThreshold: 5,
		Status:           ACTIVE,
		CrawlTimeout:     20,
		URL:              "https://case1.com",
	}
	checkURLUptime(&urlInfo1)
	assert.Equal(t, urlInfo1.FailureCount, 0)
	assert.Equal(t, urlInfo1.Status, ACTIVE)

	// Case 2: status_code != 200 AND crawl_timeout > delay
	urlInfo2 := database.URLData{
		ID:               "Case2",
		FailureCount:     0,
		FailureThreshold: 5,
		Status:           ACTIVE,
		CrawlTimeout:     20,
		URL:              "https://case2.com",
	}
	checkURLUptime(&urlInfo2)
	assert.Equal(t, urlInfo2.FailureCount, 1)
	assert.Equal(t, urlInfo2.Status, ACTIVE)

	// Case 3: crawl_timeout < delay
	urlInfo3 := database.URLData{
		ID:               "Case3",
		FailureCount:     0,
		FailureThreshold: 5,
		Status:           ACTIVE,
		CrawlTimeout:     20,
		URL:              "https://case3.com",
	}
	checkURLUptime(&urlInfo3)
	assert.Equal(t, urlInfo3.FailureCount, 1)
	assert.Equal(t, urlInfo3.Status, ACTIVE)
}
