package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	// Create backup folder
	os.MkdirAll(Configs.Folder, os.ModePerm)
	for _, filename := range Configs.Source {
		fmt.Printf("Processing %v\n", filename)
		backupName := strings.Replace(filename, filepath.Ext(filename), fmt.Sprintf("_%v%v", currentTime.Format(Configs.Format), filepath.Ext(filename)), 1)
		if Configs.OnlyChanges {
			// Get source stats
			stat, err := os.Stat(filename)
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
			if last.Size != stat.Size() || checksum(last.Name) != checksum(filename) {
				utils.Backup(filename, filepath.Join(Configs.Folder, backupName))
			}
		} else {
			utils.Backup(filename, filepath.Join(Configs.Folder, backupName))
		}
	}
}