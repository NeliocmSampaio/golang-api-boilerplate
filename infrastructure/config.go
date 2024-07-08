package infrastructure

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Net      string
	Host     string
	Port     int
	DBName   string
}

func NewConfig(filename string) (Config, error) {
	var cfg Config
	configFile, err := os.Open("./" + filename)
	defer configFile.Close()

	if err != nil {
		return cfg, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&cfg)

	return cfg, nil
}
