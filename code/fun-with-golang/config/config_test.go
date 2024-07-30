package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_GivenEnvironmentNotSetReturnDevelopmentEnv(t *testing.T) {
	expectedEnv := "development"
	actualEnv := getAppEnv()
	assert.Equal(t, expectedEnv, actualEnv)
}

func Test_GivenEnvironmentSetReturnTheEnv(t *testing.T) {
	//Set the expectedEnv
	expectedEnv := "prod"
	_ = os.Setenv("APP_ENV", expectedEnv)
	actualEnv := getAppEnv()
	assert.Equal(t, expectedEnv, actualEnv)
}

func Test_GivenAppConfigExitsByEnvReturnTrue(t *testing.T) {
	validConfigName := "development"
	dirPath := "./__resources/envs"

	viperInstance := viper.New()
	viperInstance.AddConfigPath(dirPath)
	viperInstance.SetConfigType("yml")
	viperInstance.SetConfigName(validConfigName)

	exists, _ := IsConfigFileExists(viperInstance)
	assert.True(t, exists)
	viper.Reset()
}

func Test_GivenAppConfigNotExitsByEnvReturnFalse(t *testing.T) {
	configName := "invalid-file-name"
	dirPath := "./__resources/envs"

	viperInstance := viper.New()
	viperInstance.AddConfigPath(dirPath)
	viperInstance.SetConfigType("yml")
	viperInstance.SetConfigName(configName)

	exists, _ := IsConfigFileExists(viperInstance)
	assert.False(t, exists)
	viper.Reset()
}
