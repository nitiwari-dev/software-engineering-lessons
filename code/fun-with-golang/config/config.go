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
	configType := "yml"
	env := getAppEnv()

	viperInstance := initViper(getConfigDirPaths(), configType, env)

	// Attempt to check config exists
	_, err := isConfigFileExists(viperInstance)
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

func isConfigFileExists(instance *viper.Viper) (bool, error) {
	if err := instance.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return false, err
		}
	}

	return true, nil
}

func initViper(paths []string, configType string, env string) *viper.Viper {
	viperInstance := viper.New()

	for _, path := range paths {
		viperInstance.AddConfigPath(path)
	}
	viperInstance.SetConfigName(env)
	viperInstance.SetConfigType(configType)
	viperInstance.AutomaticEnv()

	return viperInstance
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "development"
	}

	return env
}

func getConfigDirPaths() []string {
	return []string{
		"./__resources/envs",  // for Main.go
		"../__resources/envs", // for UTs
		"/etc/app/",           // for docker image
	}
}
