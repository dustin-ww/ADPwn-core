package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	adpwnConfigName   = "adpwn"
	adpwnConfigPrefix = ".adpwn"
)

func initConfig() {
	viper.AddConfigPath(moduleConfigPath)
	viper.SetConfigName(adpwnConfigName)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error while parsing adpwn configuration: %w", err))
	}
}

func RestPort() string {
	initConfig()
	if os.Getenv("REST_PORT") != "" {
		log.Println("Using REST port: " + os.Getenv("REST_PORT"))
		return os.Getenv("REST_PORT")
	}
	log.Printf("-----------" + viper.GetString(adpwnConfigPrefix+".rest_port"))
	return viper.GetString(adpwnConfigPrefix + ".rest_port")
}

func SSEPort() string {
	initConfig()
	if os.Getenv("SSE_PORT") != "" {
		return os.Getenv("SSE_PORT")
	}
	return viper.GetString(adpwnConfigPrefix + ".sse_port")
}
