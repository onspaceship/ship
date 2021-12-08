package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/onspaceship/ship/pkg/token"

	"github.com/spf13/viper"
)

const (
	defaultCoreBaseURL = "https://core.onspaceship.com/"
)

type Client struct {
	http    *http.Client
	baseURL *url.URL
	token   string
}

func NewClient() *Client {
	baseURL, err := url.Parse(viper.GetString("core_base_url"))
	if err != nil {
		log.Fatalf("Invalid Core base URL: %s", viper.GetString("core_base_url"))
	}

	http := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Client{
		http:    http,
		baseURL: baseURL,
		token:   token.GetToken(),
	}
}

func (client *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.token))
	req.Close = true

	return client.http.Do(req)
}

func (client *Client) GetJSON(url string, data interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

func (client *Client) corePath(path string, tokens ...interface{}) string {
	url, _ := client.baseURL.Parse(fmt.Sprintf(path, tokens...))
	return url.String()
}

func init() {
	viper.SetDefault("core_base_url", defaultCoreBaseURL)
}
