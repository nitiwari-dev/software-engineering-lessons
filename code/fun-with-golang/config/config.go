package config

import (
	"errors"
	"fun-with-golang/config/model"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func InitAppConfig() (*model.AppConfig, error) {

	// Init Viper Instance
	dirPath := "./resources/envs"
	configType := "yml"
	viperInstance := initViper(dirPath, configType)

	// Attempt to check config exists
	_, err := IsConfigFileExists(viperInstance)
	if err != nil {
		log.Error().Msgf("config file not found on path: %v", err)
		return nil, err
	}

	// Unmarshall if the config exits
	var appConfig model.AppConfig
	if err := viperInstance.Unmarshal(&appConfig); err != nil {
		log.Error().Msgf("erro unmarshalling the config: %v", err)
		return nil, err
	}

	return &appConfig, nil

}

func IsConfigFileExists(instance *viper.Viper) (bool, error) {
	if err := instance.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return false, err
		}
	}

	return true, nil
}

func initViper(dirPath, configType string) *viper.Viper {
	viperInstance := viper.New()
	configFileName := getAppEnv()
	viperInstance.AddConfigPath(dirPath)
	viperInstance.SetConfigName(configFileName)
	viperInstance.SetConfigType(configType)

	return viperInstance
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "development"
	}

	return env
}
