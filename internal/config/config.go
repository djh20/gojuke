package config

import (
	"flag"
)

var DebugMode bool
var ListenAddress string
var CustomDataDirectory string

func Init() {
	flag.BoolVar(&DebugMode, "d", false, "debug mode")
	flag.StringVar(&ListenAddress, "bind", ":80", "listen address")
	flag.StringVar(&CustomDataDirectory, "data-dir", "", "data directory")
	flag.Parse()
}
