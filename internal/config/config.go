package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string
}
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"development"`
	storagePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_Server"`
}

func MustLoad() *Config {
	// Load the configuration from the environment variables and YAML file
	var configpath string
	configpath = os.Getenv("CONFIG_PATH")
	if configpath == "" {
		flags := flag.String("config", "", "Path to the config file")
		flag.Parse()
		configpath = *flags
		if configpath == "" {

			log.Fatal("Config path is not set")
		}
	}
	// stat return file info and error
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configpath)
	}
	var cfg Config
	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file : %s", err.Error())
	}

	return &cfg
}
