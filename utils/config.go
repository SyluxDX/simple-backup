package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)
// Configurations json
type Configurations struct {
	Source []string `json:"source"`
	Folder string `json:"backupFolder"`
	LogFolder string `json:"logFolder"`
	Format string `json:"backupFormat"`
	OnlyChanges bool `json:"backupOnlyChanges"`
}

// GetConfigs read and parse configurations json file
func GetConfigs() (Configurations, error) {
	// json data
	var config Configurations

	// read file
	fdata, err := ioutil.ReadFile(path.Join(".", "config.json"))
	if err != nil {
		return config, err
	}

	// unmarshall it
	err = json.Unmarshal(fdata, &config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config, nil
}