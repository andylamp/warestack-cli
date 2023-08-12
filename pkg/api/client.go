package api

import (
	"net/http"
	"time"
	"warestack-cli-v2/pkg/auth"
)

type Client struct {
	HTTPClient  *http.Client
	BaseURL     string
	TokenClaims *auth.TokenClaims
}

func NewClient(tokenClaim *auth.TokenClaims) *Client {
	return &Client{
		HTTPClient:  &http.Client{Timeout: 10 * time.Second},
		BaseURL:     "https://api-dev.warestack.cloud/v1",
		TokenClaims: tokenClaim,
	}
}
