package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type GitlabSecrets struct {
	ApiToken string `json:"API_TOKEN"`
}

func LoadSecrets(filePath string) (*GitlabSecrets, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)

	var secrets GitlabSecrets

	err = json.Unmarshal(fileBytes, &secrets)
	if err != nil {
		return nil, err
	}

	return &secrets, nil
}
