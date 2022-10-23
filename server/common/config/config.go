package config

import (
	"fmt"
	"path"
	"time"
)

var defaultConfig *Config

type Servers struct {
	Host string
	Port int
}

type Default struct {
	AccessLogPath string
	LogPath       string
	Debug         bool
}

type Database struct {
	Driver       string
	Dsn          string
	MaxIdleConn  int
	MaxOpenConns int
	MaxLifeTime  time.Duration
}

type Config struct {
	Default  Default
	Servers  Servers
	Database Database
}

func InitDefaultConfig(config *Config) {
	defaultConfig = config
}

func GetConfig() *Config {
	if defaultConfig == nil {
		defaultConfig = NewDefaultConfig()
	}
	return defaultConfig
}

func NewDefaultConfig() *Config {
	config := Config{
		Default: Default{
			AccessLogPath: "logs/beluga_access.log",
			LogPath:       "logs/beluga.log",
		},
		Servers: Servers{
			Host: "127.0.0.1",
			Port: 8080,
		},
		Database: Database{
			Driver:       "sqlite",
			Dsn:          "db/beluga.db",
			MaxIdleConn:  10,
			MaxOpenConns: 100,
			MaxLifeTime:  time.Hour,
		},
	}
	return &config
}

func NewConfig(configPath string) (*Config, error) {
	suffix := path.Ext(configPath)
	fmt.Println("suffix: ", suffix)
	switch suffix {
	case ".toml":
		return NewTomlConfig(configPath)
	default:
		return NewDefaultConfig(), nil
	}

}
