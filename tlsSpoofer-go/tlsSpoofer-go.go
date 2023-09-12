package main

import (
	"flag"
	. "github.com/jellyb0y/TLS-Spoofer/tlsSpoofer-go/config"
	. "github.com/jellyb0y/TLS-Spoofer/tlsSpoofer-go/wsServer"
	"log"
	"os"
	"os/signal"
)

func main() {
	var helpFlag, EnvConfigFlag, FileConfigFlag bool
	flag.BoolVar(&EnvConfigFlag, "env-config", false, "use config from ENV")
	flag.BoolVar(&FileConfigFlag, "file-config", false, "use config from .toml file")
	flag.BoolVar(&helpFlag, "help", false, "print help")
	flag.Parse()

	if helpFlag {
		flag.Usage()
		os.Exit(0)
	} else {
		config := NewConfig(ConfigOptions{
			EnvConfig:  EnvConfigFlag,
			FileConfig: FileConfigFlag,
		})
		errCh := make(chan error)
		sigCh := make(chan os.Signal, 1)
		for {
			wsServer := NewWSServer(config)
			go func() {
				errCh <- wsServer.Serve()
			}()
			signal.Notify(sigCh, os.Interrupt)
			select {
			case err := <-errCh:
				wsServer.FinishWork()
				log.Printf("failed to serve: %v\n\ntrying to restart server...", err)
			case sig := <-sigCh:
				log.Printf("terminating: %v", sig)
				wsServer.FinishWork()
				os.Exit(0)
			}
		}
	}
}
