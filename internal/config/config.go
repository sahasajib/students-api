package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env        string `yaml:"env" env:"ENV" envDefault:"dev" env-required:"true"`
	StoragePath    string `yaml:"storage" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}


func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path must be provided either through the CONFIG_PATH environment variable or the --config flag")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", configPath)
	}

	var cfg Config 
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read configuration file: %v", err)
	}
	return &cfg
}