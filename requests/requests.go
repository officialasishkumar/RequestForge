package requests

import (
	"RequestForge/logger"
	"RequestForge/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadRequests(jsonFilePath string) ([]models.Request, error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		logger.Logger.Println("Error opening JSON file:", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var requests []models.Request
	err = json.Unmarshal(byteValue, &requests)
	if err != nil {
		fmt.Println("Error parsing JSON file:", err)
		logger.Logger.Println("Error parsing JSON file:", err)
		return nil, err
	}

	return requests, nil
}
