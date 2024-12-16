package storage

import (
	"log"
	"os"
	"path/filepath"

	"github.com/djh20/openjukebox/internal/config"
)

var DataDirectory string

func Init() {
	DataDirectory = config.CustomDataDirectory

	if DataDirectory == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		DataDirectory = filepath.Join(homeDir, "openjukebox")
	}

	err := os.Mkdir(DataDirectory, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	log.Println("Data directory:", DataDirectory)
}

func GetDataSubDirectory(relativePath string) string {
	absolutePath := filepath.Join(DataDirectory, relativePath)

	err := os.MkdirAll(absolutePath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	return absolutePath
}
