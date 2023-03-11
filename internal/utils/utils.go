package utils

import (
	"cashier-api/core/logger"
	"encoding/json"
	"os"
)

// ReadJSONFile read json file
func ReadJSONFile(path string, entities interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		logger.Logger.Errorf("[ReadJSONFile] read file error: %s", err)
		return err
	}

	err = json.Unmarshal([]byte(data), entities)
	if err != nil {
		logger.Logger.Errorf("[ReadJSONFile] Unmarshal error: %s", err)
		return err
	}

	return nil
}

// WriteJsonFile write json file
func WriteJsonFile(path string, entities interface{}) error {
	data, err := json.Marshal(entities)
	if err != nil {
		logger.Logger.Errorf("[WriteJsonFile] Marshal error: %s", err)
		return err
	}

	err = os.WriteFile(path, data, 0666)
	if err != nil {
		logger.Logger.Errorf("[WriteJsonFile] write file error: %s", err)
		return err
	}

	return nil
}
