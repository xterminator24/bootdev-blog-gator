package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	// set c.CurrentUserName
	c.CurrentUserName = name
	return write(c)
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	// json.Unmarshal into Config
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	// locate home dir
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// build full path to ~/.gatorconfig.json
	configFile := filepath.Join(home, configFileName)

	return configFile, nil
}

func write(cfg *Config) error {
	// write config back to config
	data, err := json.MarshalIndent(cfg, "", "    ") // pretty print
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
