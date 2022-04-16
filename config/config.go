package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

var (
	once sync.Once
	pc   projectConfig
	rc   redisConfig
)

type projectConfig struct {
	Title string      `toml:"title"`
	Redis redisConfig `toml:"redis"`
}

type redisConfig struct {
	Enabled  bool   `toml:"enabled"`
	Address  string `toml:"address"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
}

func loadConfig() {
	once.Do(func() {
		configPath := "config.toml"
		configContent, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("os.ReadFile出错", err)
		}
		err = toml.Unmarshal(configContent, &pc)
		if err != nil {
			fmt.Println("UnmarshalConfig出错", err)
		}
		rc = pc.Redis
	})
}

// GetConfig: 获取项目配置结构体
func GetConfig() projectConfig {
	loadConfig()
	return pc
}

// GetRedisConfig: 获取redis配置结构体
func GetRedisConfig() redisConfig {
	loadConfig()
	return rc
}
