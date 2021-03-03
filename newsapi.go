package gonewsapi

import (
	"fmt"
)

var baseURl = "http://newsapi.org/v2/"
var API_KEYS = "45537dd2a5f342879da141829ce12615"

func (n *NewsApiClient) GetHeadlines(category string, country string) (*Response, error) {
	getHeadlines := &Response{}
	endpoint := fmt.Sprintf(baseURl + "top-headlines?country=" + country + "&category=" + category + "&apiKey=" + API_KEYS)
	err := n.makeRequest("GET", endpoint, nil, nil, getHeadlines)
	if err != nil {
		return nil, err
	}

	return getHeadlines, nil
}
