package utils

import (
    "fmt"
    "path"
    "encoding/json"
    "io/ioutil"
)
// Configurations json
type Configurations struct {
    Source string `json:"source"`
    Folder string `json:"backupFolder"`
	Format string `json:"backupFormat"`
	OnlyChanges bool `json:"backupOnlyChanges"`
}

// GetConfigs read and parse configurations json file
func GetConfigs() Configurations {
        // read file
        fdata, err := ioutil.ReadFile(path.Join(".", "config.json"))
        if err != nil {
          fmt.Print(err)
        }
        // json data
        var config Configurations
        // unmarshall it
        err = json.Unmarshal(fdata, &config)
        if err != nil {
            fmt.Println("error:", err)
        }
        return config
}