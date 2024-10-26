package main

import (
	"log"

	"github.com/djh20/gojuke/internal/config"
	"github.com/djh20/gojuke/internal/music"
	"github.com/djh20/gojuke/internal/storage"
	"github.com/djh20/gojuke/internal/web"
)

func main() {
	log.Println("Initializing application")
	storage.Init()
	config.Init()
	music.Init()
	web.Init()

	go music.Run()
	web.Run()
}
