package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ReadJSON reads a JSON file and decodes it into the given interface.
func ReadJSON(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	if err := json.Unmarshal(data, v); err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

// WriteJSON writes the given interface as JSON to the given file.
func WriteJSON(filePath string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

// UpdateJSONField updates a specific field in a JSON file.
func UpdateJSONField(filePath string, key string, value interface{}) error {
	// Load existing JSON data from the file.
	var jsonData map[string]interface{}
	err := ReadJSON(filePath, &jsonData)
	if err != nil {
		return err
	}

	// Update the desired key-value.
	jsonData[key] = value

	// Write the modified data back to the file.
	return WriteJSON(filePath, jsonData)
}
