package config

import "github.com/BurntSushi/toml"

func NewTomlConfig(configPath string) (*Config, error) {
	config := NewDefaultConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
