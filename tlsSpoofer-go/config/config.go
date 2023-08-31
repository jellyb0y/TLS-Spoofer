package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Options struct {
	EnvConfig  bool
	FileConfig bool
}

// Config конфиг WSSserver
type Config struct {
	ListenPort      int `toml:"listen-port"`
	ReadTimeout     int `toml:"read-timeout"`
	GoroutinesCount int `toml:"goroutines-count"`
	ReadBufferSize  int `toml:"read-buffer-size"`
	WriteBufferSize int `toml:"write-buffer-size"`
}

// NewConfig конструктор Config с дефолтными значениями
func NewConfig(options Options) *Config {

	// настройки по умолчанию
	config := &Config{
		ListenPort:      8080,
		ReadTimeout:     1,
		GoroutinesCount: 10,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	switch {
	// взять настройки из ENV
	case options.EnvConfig:
		{
			listenPort, existsListenPort := os.LookupEnv("WS_LISTEN_PORT")
			if existsListenPort {
				listenPortValue, err := strconv.Atoi(listenPort)
				if err != nil {
					log.Fatal(err)
				}
				config.ListenPort = listenPortValue
			}
			readTimeout, existsReadTimeout := os.LookupEnv("WS_READ_TIMEOUT")
			if existsReadTimeout {
				readTimeoutValue, err := strconv.Atoi(readTimeout)
				if err != nil {
					log.Fatal(err)
				}
				config.ReadTimeout = readTimeoutValue
			}
			goroutinesCount, existsMaxConcurrency := os.LookupEnv("WS_GOROUTINES_COUNT")
			if existsMaxConcurrency {
				goroutinesCountValue, err := strconv.Atoi(goroutinesCount)
				if err != nil {
					log.Fatal(err)
				}
				config.GoroutinesCount = goroutinesCountValue
			}
			readBufferSize, existsReadBufferSize := os.LookupEnv("WS_READ_BUFFER_SIZE")
			if existsReadBufferSize {
				readBufferSizeValue, err := strconv.Atoi(readBufferSize)
				if err != nil {
					log.Fatal(err)
				}
				config.ReadTimeout = readBufferSizeValue
			}
			writeBufferSize, existsWriteBufferSize := os.LookupEnv("WS_READ_BUFFER_SIZE")
			if existsWriteBufferSize {
				writeBufferSizeValue, err := strconv.Atoi(writeBufferSize)
				if err != nil {
					log.Fatal(err)
				}
				config.ReadTimeout = writeBufferSizeValue
			}
			log.Println("using ENV config")

			return config
		}
	// пропарсить настройки из файла конфигурации
	case options.FileConfig:
		{
			name := filepath.Base(os.Args[0])
			configFile := fmt.Sprintf("../etc/conf.d/%s.toml", name)

			configBodyBytes, err := ioutil.ReadFile(configFile)
			if err != nil {
				flag.Usage()
				log.Fatal(err)
			}

			if _, err = toml.Decode(string(configBodyBytes), config); err != nil {
				log.Fatal(err)
			}
			log.Println("using config from .toml file")

			return config
		}
	default:
		log.Println("using default config")
		return config
	}
}
