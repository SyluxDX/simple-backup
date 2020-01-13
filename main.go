package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"strings"
	"./utils"
)

// Global Variables
var (
	Configs utils.Configurations
)

func checksum(name string) (string) {
	source, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer source.Close()

	hash := md5.New()
	_, err = io.Copy(hash, source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func main() {
	Configs = utils.GetConfigs()
	currentTime := time.Now()
	backupName := strings.Replace(Configs.Source, filepath.Ext(Configs.Source), fmt.Sprintf("_%v%v", currentTime.Format(Configs.Format), filepath.Ext(Configs.Source)), 1)

	// Create backup folder
	os.MkdirAll(Configs.Folder, os.ModePerm)
	
	if Configs.OnlyChanges {
		// Get source stats
		stat, err := os.Stat(Configs.Source)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		// Get last backup stats
		last, err := utils.GetLastBackup(Configs.Folder)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		if last.Size != stat.Size() || checksum(last.Name) != checksum(Configs.Source) {
			utils.Backup(Configs.Source, filepath.Join(Configs.Folder, backupName))
		}
	} else {
		utils.Backup(Configs.Source, filepath.Join(Configs.Folder, backupName))
	}
}