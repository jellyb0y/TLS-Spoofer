package config

// Config конфиг WSSserver
type Config struct {
	ListenPort     int `toml:"listen-port"`
	ReadTimeout    int `toml:"read-timeout"`
	MaxConcurrency int `toml:"max-concurency"`
}

// NewConfig конструктор Config с дефолтными значениями
func NewConfig() *Config {
	return &Config{
		ListenPort:     8080,
		ReadTimeout:    1,
		MaxConcurrency: 10,
	}
}
