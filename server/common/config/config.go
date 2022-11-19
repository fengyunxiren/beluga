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

type Auth struct {
	SecretKey   string
	ExpireTime  time.Duration
	RefreshTime time.Duration
	Issuer      string
}

type Config struct {
	Default  Default
	Servers  Servers
	Database Database
	Auth     Auth
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
		Auth: Auth{
			SecretKey:   `e05b5334052af5d521856506a2fcf5d9c9f3a82bf5e9fdae0c5e8e7f22f3e000201b50a87df1b3ca57f77f7d30530204b9614aaac1cdfb01d8780fb67caff7a4`,
			ExpireTime:  time.Hour,
			RefreshTime: time.Hour * 24,
			Issuer:      `beluga`,
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
