package api

import (
	"encoding/json"
	"fmt"
	"github.com/warestack/warestack-core-api/modules/dto"
	"net/http"
)

type UserResponse struct {
	Message string          `json:"message"`
	User    *dto.UserOutput `json:"user"`
}

func (c *Client) GetUserProfile() (*dto.UserOutput, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/me", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Credentials.IDToken))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var userResponse *UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResponse); err != nil {
		return nil, err
	}

	return userResponse.User, nil
}
