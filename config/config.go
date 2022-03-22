package config

import (
	"PowerShare/helper"
	"embed"
	"log"
)

//go:embed *.env.default
var EnvFiles embed.FS

// GetConfigFilePath returns the path where the env files are expected.
// If it does not exist the empty folder will be created.
func GetConfigFilePath() string {
	const configDirPath = "config"
	if !helper.Exists(configDirPath) {
		err := helper.CreateDir(configDirPath)
		if err != nil{
			log.Println(err)
		}
	}
	return configDirPath
}
