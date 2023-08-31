package main

import (
	"flag"
	"github.com/jellyb0y/TLS-Spoofer/tlsSpoofer-go/config"
	. "github.com/jellyb0y/TLS-Spoofer/tlsSpoofer-go/wsServer"
	"log"
	"os"
)

var EnvConfigFlag, FileConfigFlag bool

func main() {
	var helpFlag bool
	flag.BoolVar(&EnvConfigFlag, "env-config", false, "use config from ENV")
	flag.BoolVar(&FileConfigFlag, "file-config", false, "use config from .toml file")
	flag.BoolVar(&helpFlag, "help", false, "print help")
	flag.Parse()
	if helpFlag {
		flag.Usage()
		os.Exit(0)
	} else {
		wsServer := NewWSServer(config.Options{
			EnvConfig:  EnvConfigFlag,
			FileConfig: FileConfigFlag,
		})
		log.Fatal(wsServer.Serve())
	}
}
