package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"./utils"
)

// Global Variables
var (
	Configs utils.Configurations
	Log *log.Logger
)

func checksum(name string) (string) {
	source, err := os.Open(name)
	if err != nil {
		Log.Fatalln(err)
	}
	defer source.Close()

	hash := md5.New()
	_, err = io.Copy(hash, source)
	if err != nil {
		Log.Fatalln(err)
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func main() {
	currentTime := time.Now()
	Configs, err := utils.GetConfigs()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(Configs.Folder)
	// Create backup and log folders
	os.MkdirAll(Configs.Folder, os.ModePerm)
	os.MkdirAll(Configs.LogFolder, os.ModePerm)
	// Create Logger
	logFile, err := os.OpenFile(filepath.Join(Configs.LogFolder, fmt.Sprintf("backup_%v.log", currentTime.Format("20060102"))), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Log.Fatalln(err)
	}
	defer logFile.Close()
	Log = log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)

	Log.Println("Starting Backup Process")
	for _, filename := range Configs.Files {
		Log.Printf("Processing %v\n", filename)
		backupName := strings.Replace(filename, filepath.Ext(filename), fmt.Sprintf("_%v%v", currentTime.Format(Configs.Format), filepath.Ext(filename)), 1)
		if Configs.OnlyChanges {
			// Get source stats
			stat, err := os.Stat(filepath.Join(Configs.Source, filename))
			if err != nil {
				Log.Fatalln(err)
			}
			// Get last backup stats
			last, err := utils.GetLastBackup(Configs.Folder)
			if err != nil {
				Log.Fatalln(err)
			}
			if last.Size != stat.Size() || checksum(last.Name) != checksum(filepath.Join(Configs.Source, filename)) {
				out, err := utils.Backup(filepath.Join(Configs.Source, filename), filepath.Join(Configs.Folder, backupName))
				if err != nil {
					Log.Fatalln(err)
				}
				Log.Println(out)
			}
		} else {
			out, err := utils.Backup(filepath.Join(Configs.Source, filename), filepath.Join(Configs.Folder, backupName))
			if err != nil {
				Log.Fatalln(err)
			}
			Log.Println(out)
		}
	}
}