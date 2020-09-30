package httpRequests

import (
	"github.com/gojektech/heimdall/httpclient"
	"net/http"
	"time"
)

//func init() {
//	SetHTTPController(&MonitorHttp{})
//}

type HttpController interface {
	MakeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error)
}

type MonitorHttp struct{}

var httpCalls HttpController

func SetHTTPController(hType HttpController) {
	httpCalls = hType
}

func GetHTTPController() HttpController {
	return httpCalls
}

func (mh *MonitorHttp) MakeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error) {

	timeout := time.Duration(crawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	return client.Get(url, nil)
}
