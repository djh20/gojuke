package config

import (
	"flag"
)

var DebugMode bool
var ListenAddress string

func Init() {
	flag.BoolVar(&DebugMode, "d", false, "debug mode")
	flag.StringVar(&ListenAddress, "a", ":80", "listen address")
	flag.Parse()
}
