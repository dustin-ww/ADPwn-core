package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	configLoaded        sync.Once
	modulesConfigLoaded sync.Once
)

func initViper() {
	viper.SetConfigName(adpwnConfigName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(moduleConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading main config: %v", err)
	}
}

func initModulesViper() {
	viper.SetConfigName(moduleConfigName)
	viper.AddConfigPath(moduleConfigPath)

	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error merging modulelib config: %v", err)
	}
}

func Init() {
	configLoaded.Do(initViper)
}

func InitModules() {
	modulesConfigLoaded.Do(initModulesViper)
}
