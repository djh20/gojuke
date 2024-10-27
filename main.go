package main

import (
	"embed"
	"log"

	"github.com/djh20/openjukebox/internal/config"
	"github.com/djh20/openjukebox/internal/music"
	"github.com/djh20/openjukebox/internal/storage"
	"github.com/djh20/openjukebox/internal/web"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {
	log.Println("Initializing application")
	storage.Init()
	config.Init()
	music.Init()
	web.Init(frontend)

	go music.Run()
	web.Run()
}
