package auth

import (
	"errors"
)

// TokenClaims captures the structure of the access and refresh tokens.
type TokenClaims struct {
	IdToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
}

// EnsureValidToken checks if the access token is valid, if not, it tries to refresh it.
// If the refresh token is also invalid, it returns an error.
func EnsureValidToken(tokenClaim *TokenClaims) error {
	if tokenClaim == nil {
		return errors.New("no tokens provided")
	}

	// TODO: Check access token validity
	// If access token is valid, return

	// TODO: Check refresh access token validity
	// If refresh token is also invalid, return an error

	// TODO: Refresh the access token using the refresh token.
	// Add your code here to make a request to your server to refresh the access token.

	return nil
}
