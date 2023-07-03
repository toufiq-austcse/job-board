package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	defaultPort                = "3000"
	defaultAppName             = "Boilerplate"
	defaultDBDebugEnabledValue = "false"
)

type (
	Config struct {
		PORT           string
		APP_NAME       string
		APP_URL        string
		DB_DRIVER_NAME string
		DB_CONFIG      DB_CONFIG
	}
	DB_CONFIG struct {
		DB_NAME       string
		HOST          string
		PORT          string
		USER          string
		PASSWORD      string
		DEBUG_ENABLED string
	}
)

var AppConfig Config

func Init(envFilePath string) {
	populateDefault()
	parseConfigFile(envFilePath, "main")
	setFromEnv(&AppConfig)

	fmt.Println("Configuration loaded")

}

func setFromEnv(config *Config) {
	config.PORT = viper.GetString("PORT")
	config.APP_NAME = viper.GetString("APP_NAME")
	config.APP_URL = viper.GetString("APP_URL")
	config.DB_CONFIG.DB_NAME = viper.GetString("DB_NAME")
	config.DB_CONFIG.HOST = viper.GetString("DB_HOST")
	config.DB_CONFIG.PORT = viper.GetString("DB_PORT")
	config.DB_CONFIG.PASSWORD = viper.GetString("DB_PASSWORD")
	config.DB_DRIVER_NAME = viper.GetString("DB_DRIVER_NAME")
	config.DB_CONFIG.USER = viper.GetString("DB_USER")
	config.DB_CONFIG.DEBUG_ENABLED = viper.GetString("DB_DEBUG_ENABLED")
}

func parseConfigFile(envFilePath, configName string) {
	viper.SetConfigType("env")
	viper.SetConfigName(configName)
	viper.SetConfigFile(envFilePath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error in reading config ", err)
	}
	viper.AutomaticEnv()

	err := viper.MergeInConfig()
	if err != nil {
		fmt.Println("err in merge config", err)
	}
}
func populateDefault() {
	viper.SetDefault("PORT", defaultPort)
	viper.SetDefault("APP_NAME", defaultAppName)
	viper.SetDefault("DEBUG_ENABLED", defaultDBDebugEnabledValue)
}
