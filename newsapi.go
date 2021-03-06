package gonewsapi

import (
	"fmt"
	"net/url"
)

var baseURL = "http://newsapi.org/v2/"
var API_KEYS = "45537dd2a5f342879da141829ce12615"

func BuildUrl(base string, args ...string) (endpoint string) {
	var params string
	for i, param := range args {
		params += param
		if i < len(args)-1 {
			params += "&"
		}
	}
	p := url.PathEscape(params)
	u := base + p

	return u
}

func (n *NewsApiClient) GetHeadlines(args []string) (*Response, error) {
	if len(args) == 0 {
		return nil, ArgumentErr
	}

	getHeadlines := &Response{}
	endpoint := BuildUrl(fmt.Sprintf("%s/%s", baseURL, "top-headlines"), args...)
	err := n.makeRequest("GET", endpoint, nil, nil, getHeadlines)
	if err != nil {
		return nil, err
	}

	return getHeadlines, nil
}

func (n *NewsApiClient) GetEverything(args []string) (*Response, error) {
	if len(args) == 0 {
		return nil, ArgumentErr
	}
	getEverything := &Response{}
	endpoint := BuildUrl(fmt.Sprintf("%s/%s", baseURL, "everything"), args...)
	err := n.makeRequest("GET", endpoint, nil, nil, getEverything)
	if err != nil {
		return nil, err
	}

	return getEverything, nil
}
