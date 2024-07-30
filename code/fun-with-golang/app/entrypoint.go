package app

import (
	"fun-with-golang/config"
	"fun-with-golang/intro"
	"github.com/rs/zerolog/log"
)

func StartApp() {
	appConfig, err := config.InitAppConfig()
	if err != nil {
		log.Err(err).Msgf("failed to init the app config")
	}

	//Check to enable intro
	if appConfig.ServiceConfig.IntroConfig.IsEnabled {
		intro.Intro()
	}
}
