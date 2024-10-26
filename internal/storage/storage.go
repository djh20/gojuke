package storage

import (
	"log"
	"os"
	"path/filepath"
)

var DataDirectory string

func Init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	DataDirectory = filepath.Join(homeDir, "gojuke")

	err = os.Mkdir(DataDirectory, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	log.Println("Found data directory:", DataDirectory)
}

func GetDataSubDirectory(relativePath string) string {
	absolutePath := filepath.Join(DataDirectory, relativePath)

	err := os.MkdirAll(absolutePath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return absolutePath
}
