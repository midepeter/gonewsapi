package gonewsapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var API_KEY = "45537dd2a5f342879da141829ce12615"

type (
	Error struct {
		Code     int
		Body     string
		Endpoint string
	}

	Header struct {
		key   string
		Value string
	}

	NewsApiClient struct {
		client *http.Client
	}
)

func New(cli *http.Client) (*NewsApiClient, error) {
	return &NewsApiClient{
		client: cli,
	}, nil
}

func (n *NewsApiClient) makeRequest(method, url string, body io.Reader, headers []Header, responseTarget interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	for _, h := range headers {
		req.Header.Set(h.key, h.Value)
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := n.client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		err := json.Unmarshal(b, responseTarget)
		if err != nil {
			return err
		}
		return nil
	}

	return err
}
