package api

import (
	"encoding/json"
	"fmt"
	"github.com/warestack/warestack-core-api/modules/dto"
	"net/http"
)

type OrganizationResponse struct {
	//Message      string          `json:"message"`
	Organization *dto.OrganizationOutput
}

func (c *Client) GetOrganizationByID(id uint) (*dto.OrganizationOutput, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations/%d", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.TokenClaims.IdToken))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var organizationResponse *dto.OrganizationOutput
	if err := json.NewDecoder(resp.Body).Decode(&organizationResponse); err != nil {
		return nil, err
	}
	fmt.Println(organizationResponse.Alias)
	return organizationResponse, nil
}
