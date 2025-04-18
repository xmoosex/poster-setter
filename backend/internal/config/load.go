package config

import (
	"fmt"
	"os"
	"path"
	"poster-setter/internal/modals"

	"gopkg.in/yaml.v3"
)

// Global is a pointer to the global configuration instance.
// It is used throughout the application to access configuration settings.
var Global *modals.Config

// LoadYamlConfig loads the application configuration from a YAML file.
//
// Steps:
// 1. Retrieve the configuration file path from the `/config/config.yml`
// 2. Check if the YAML file exists at the specified path.
// 3. Read and parse the YAML file into a `Config` struct.
// 4. Set the global `Global` variable to the loaded configuration.
//
// Returns:
//   - A pointer to the `Config` struct containing the loaded configuration.
//   - An error if the configuration file is missing, unreadable, or invalid.
func LoadYamlConfig() (*modals.Config, error) {

	// Use an environment variable to determine the config path
	// By default, it will use /config
	// This is useful for testing and local development
	// In Docker, the config path is set to /config
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/config"
	}

	// Check for a config.yml or config.yaml file
	yamlFile := path.Join(configPath, "config.yml")
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		yamlFile = path.Join(configPath, "config.yaml")
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return nil, fmt.Errorf("config.yml file not found in %s", configPath)
		}
	}

	// Read the YAML file
	data, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config.yml file")
	}

	// Parse the YAML file into a Config struct
	var config modals.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config.yml file")
	}

	Global = &config
	return &config, nil
}
