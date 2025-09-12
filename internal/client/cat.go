package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
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
	u, err := url.Parse(baseURL + "/breeds/search")
	if err != nil {
		return false, fmt.Errorf("parse base url: %w", err)
	}
	q := u.Query()
	q.Set("q", breed)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return false, fmt.Errorf("build request: %w", err)
	}
	if c.apiKey != "" {
		req.Header.Set("x-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("cat api returned status: %s", resp.Status)
	}

	type cat struct {
		ID string `json:"id"`
	}

	var res []cat
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return false, errwrap.Wrap("decode response", err)
	}

	return len(res) > 0, nil
}
