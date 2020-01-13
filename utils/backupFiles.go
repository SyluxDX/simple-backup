package utils

import (
	"fmt"
	"io"
	"os"
    "path/filepath"
)
// Info Backup File info
type Info struct {
    Name string 
	Size int64
}

// Backup Backups file
func Backup(src string, dst string) {
	// Open source file
	source, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer source.Close()
	// Open destination file
	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer destination.Close()
	// copy file
	nBytes, err := io.Copy(destination, source)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Written %d bytes\n", nBytes)
}

// GetLastBackup get last backup file info
func GetLastBackup(dir string) (Info, error){
	var file Info
	var err error
	var time int64
	file.Size = 0
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // check if it is a regular file (not dir)
        if info.Mode().IsRegular() {
			if info.ModTime().Unix() > time {
				file.Name = path
				file.Size = info.Size()
				time = info.ModTime().Unix()
			}
        }
        return nil
	})
	return file, err
}