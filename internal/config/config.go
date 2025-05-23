package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type Config struct {
	Skiplist SkiplistConfig `json:"skiplist"`
	//Add other config
}

type SkiplistConfig struct {
	MaxLevel    int     `json:"max_level"`
	Probability float64 `json:"probability"`
}

// Default Configuration
func DefaultConfig() *Config {
	return &Config{
		Skiplist: SkiplistConfig{
			MaxLevel:    10,
			Probability: 0.5,
		},
	}
}

// Load JSON configuration file
func LoadConfig(filename string) (*Config, error) {
	Cfg := DefaultConfig() //Fall back to default config if error

	file, err := os.Open(filename)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			//If config file not found create a new one
			fmt.Println("Configuration file not found, generating a new file")
			if err := generateDefaultConfig(filename, Cfg); err != nil {
				fmt.Printf("Error generating default configuration: %v", err)
			}
			fmt.Printf("default configuration file '%s' generated", filename)
			return Cfg, nil
		}
		//If error opening config file
		return Cfg, fmt.Errorf("Cannot open config file at '%s'. '%w'", filename, err)
	}
	defer file.Close()
	//Unmarshal config from JSON config file
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(Cfg); err != nil {
		return Cfg, fmt.Errorf("error decoding config file '%s': %w", filename, err)
	}
	return Cfg, nil
}

// Marshal default config to JSON and save it to config file
func generateDefaultConfig(filename string, Cfg *Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating config file '%s': %w", filename, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(Cfg); err != nil {
		return fmt.Errorf("error encoding default config to '%s': %w", filename, err)
	}

	return nil
}
