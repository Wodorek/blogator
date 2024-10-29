package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	cfgPath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(cfgPath)

	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	cfg := Config{}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(userName string) error {

	cfg.CurrentUserName = userName

	return write(*cfg)

}

func write(cfg Config) error {

	fullPath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homePath, configFileName)

	return fullPath, nil

}
