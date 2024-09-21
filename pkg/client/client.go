package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

// Client responsible for making HTTP requests
type Client struct {
	httpClient *http.Client
	baseUrl    string
}

// NewClient initializes and returns a new instance of the Client struct
func NewClient(baseUrl string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseUrl,
	}
}

// GetDateTime performs an HTTP GET request to fetch the current date and time
func (c *Client) GetDateTime(isJson bool) (string, error) {
	if isJson {
		return c.getDateTimeJson()
	} else {
		return c.getDateTimePlain()
	}

}

// Retry retries an operation that returns an error using an exponential backoff strategy
func Retry(o func() error) error {
	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.InitialInterval = 500 * time.Millisecond
	expBackoff.MaxInterval = 5 * time.Second
	expBackoff.MaxElapsedTime = 30 * time.Second

	return backoff.Retry(o, expBackoff)
}

func (c *Client) getDateTimeJson() (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+"/datetime/json", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get datetime: status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response struct {
		DateTime string `json:datetime`
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return "", err
	}
	dateTime := response.DateTime

	return dateTime, nil
}

func (c *Client) getDateTimePlain() (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseUrl+"/datetime/plain", nil)
	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get datetime: status code %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	dateTime := string(data)

	return dateTime, nil
}
