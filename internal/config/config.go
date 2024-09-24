package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("%v/.gatorconfig.json", homeDir)

	return filePath, nil
}

func Read() (Config, error) {
	dbConfig := Config{}

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	dat, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(dat, &dbConfig)
	if err != nil {
		return Config{}, err
	}

	return dbConfig, nil
}

func write(cfg Config) error {
	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, dat, 0666)
	if err != nil {
		return err
	}

	return nil
}
