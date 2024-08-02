package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func Test_GivenEnvironmentIsNotSetReturnDevelopmentEnv(t *testing.T) {
	expectedEnv := "development"
	actualEnv := getAppEnv()
	assert.Equal(t, expectedEnv, actualEnv)
}

func Test_GivenEnvironmentIsSetReturnTheEnv(t *testing.T) {
	expectedEnv := "prod"
	_ = os.Setenv("APP_ENV", expectedEnv)
	actualEnv := getAppEnv()
	assert.Equal(t, expectedEnv, actualEnv)
}

func Test_GivenAppConfigFileExitsReturnTrue(t *testing.T) {
	validConfigName := "development"
	dirPath := "../__resources/envs"

	viperInstance := viper.New()
	viperInstance.AddConfigPath(dirPath)
	viperInstance.SetConfigType("yml")
	viperInstance.SetConfigName(validConfigName)

	exists, _ := isConfigFileExists(viperInstance)
	assert.True(t, exists)
}

func Test_GivenAppConfigFileDoNotExitsReturnFalse(t *testing.T) {
	configName := "invalid-file-name"
	dirPath := "../__resources/envs"

	viperInstance := viper.New()
	viperInstance.AddConfigPath(dirPath)
	viperInstance.SetConfigType(".") // add the current directly path
	viperInstance.SetConfigType("yml")
	viperInstance.SetConfigName(configName)

	exists, _ := isConfigFileExists(viperInstance)
	assert.False(t, exists)
}

func Test_GivenAppConfigFileExitsReturnConfigWithoutError(t *testing.T) {
	_ = os.Setenv("APP_ENV", "development")
	_, err := InitAppConfig()
	assert.Nil(t, err)
}

func Test_GivenAppConfigFileExitsReturnParsedConfigValue(t *testing.T) {
	_ = os.Setenv("APP_ENV", "development")
	appConfig, _ := InitAppConfig()
	assert.NotEmpty(t, appConfig.AppName)
}

func TestMain(m *testing.M) {
	log.Info().Msg("setting up the test environment")
	os.Clearenv()
	viper.Reset()

	code := m.Run()
	os.Exit(code)

}

type ViperMock interface {
	mock.Mock
}
