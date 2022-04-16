package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	myAssert := assert.New(t)

	configFilePath := "config.toml"
	backupConfigFilePath := "config.toml.bak"
	_, fileStaterr := os.Stat(configFilePath)
	if fileStaterr != nil {
		os.Rename(configFilePath, backupConfigFilePath)
	}

	configContent := []byte(`title = "hashcollision config title"

	[redis]
	enabled = true
	address = "localhost"
	port = 6379
	password = "123"`)
	err := os.WriteFile(configFilePath, configContent, 0644)
	myAssert.Nil(err)

	pc := GetConfig()

	correctProjectConfig := projectConfig{
		"hashcollision config title",
		redisConfig{
			true,
			"localhost",
			6379,
			"123",
		},
	}

	myAssert.Equal(correctProjectConfig, pc)

	os.Remove(configFilePath)
	if fileStaterr != nil {
		os.Rename(backupConfigFilePath, configFilePath)
	}
}
