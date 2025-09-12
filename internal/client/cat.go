package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://api.thecatapi.com/v1"
)

type client struct {
	apiKey     string
	httpClient *http.Client
}

func New(apiKey string) *client {
	return &client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   3 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:        10,
				IdleConnTimeout:     30 * time.Second,
				TLSHandshakeTimeout: 3 * time.Second,
			},
		},
	}
}

func (c *client) Exists(ctx context.Context, breed string) (bool, error) {
	u := fmt.Sprintf("%s/breeds/search?q=%s", baseURL, url.QueryEscape(breed))

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if c.apiKey != "" {
		req.Header.Set("x-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("cat api returned status: %d", resp.StatusCode)
	}

	type cats struct {
		ID string `json:"breed"`
	}

	var res []cats
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return false, err
	}

	return len(res) > 0, nil
}
