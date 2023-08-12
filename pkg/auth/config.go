package auth

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configPath = "config.json"

// storeTokenClaims writes the given token data to a .token.json file.
func storeTokenClaims(tokenClaims TokenClaims) error {
	data, err := json.MarshalIndent(tokenClaims, "", "")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configPath, data, 0600)
}

// ReadTokenClaims reads the token data from the .token.json file.
func ReadTokenClaims() (*TokenClaims, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read the tokens from file: %s", err)
		return nil, err
	}

	var tokenClaims TokenClaims
	if err := json.Unmarshal(data, &tokenClaims); err != nil {
		log.Fatalf("Failed to decode tokens: %s", err)
		return nil, err
	}

	return &tokenClaims, nil
}
