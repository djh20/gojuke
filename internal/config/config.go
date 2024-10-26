package config

import (
	"flag"
	"log"
)

var DebugMode bool
var ListenAddress string

func Init() {
	flag.BoolVar(&DebugMode, "d", false, "debug mode")
	flag.StringVar(&ListenAddress, "a", ":80", "listen address")
	flag.Parse()

	log.Println("Loaded configuration from command line flags")
}
