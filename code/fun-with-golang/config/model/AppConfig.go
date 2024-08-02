package model

type AppConfig struct {
	AppName          string           `mapstructure:"app_name"`
	HttpServerConfig HttpServerConfig `mapstructure:"http_server_config"`
	ServiceConfig    ServiceConfig    `mapstructure:"service_config"`
}

type HttpServerConfig struct {
	Port string `mapstructure:"port"`
}

type ServiceConfig struct {
	IntroConfig IntroConfig `mapstructure:"intro_config"`
}

type IntroConfig struct {
	IsEnabled bool `mapstructure:"is_enabled"`
}
