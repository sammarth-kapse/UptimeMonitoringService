package monitor

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"testing"
)

func TestCheckURLUptime(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepositoryController(ctrl)
	mockHttp := NewMockHttpController(ctrl)

	mockRepo.EXPECT().databaseSave(gomock.Any()).Return(nil).MaxTimes(2)
	mockRepo.EXPECT().databaseGet(gomock.Any()).Return(nil).MaxTimes(3)
	setRepoController(mockRepo)

	mockHttp.EXPECT().makeHTTPGetRequest(gomock.Any(), "https://case1.com").Return(
		&http.Response{StatusCode: http.StatusOK}, nil)

	mockHttp.EXPECT().makeHTTPGetRequest(gomock.Any(), "https://case2.com").Return(
		&http.Response{StatusCode: http.StatusConflict}, nil)

	mockHttp.EXPECT().makeHTTPGetRequest(gomock.Any(), "https://case3.com").Return(
		&http.Response{StatusCode: http.StatusConflict}, fmt.Errorf("SomeError"))

	setHTTPController(mockHttp)
	// Case 1: status_code = 200 AND crawl_timeout > delay
	urlInfo1 := URLData{
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
	urlInfo2 := URLData{
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
	urlInfo3 := URLData{
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
