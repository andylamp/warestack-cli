package api

import (
	"net/http"
	"time"
	"warestack-cli/pkg/auth"
)

type Client struct {
	HTTPClient  *http.Client
	BaseURL     string
	Credentials *auth.Credentials
}

func NewClient(credentials auth.Credentials) *Client {
	return &Client{
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
		BaseURL:     "https://api.warestack.com/v1",
		Credentials: &credentials,
	}
}
