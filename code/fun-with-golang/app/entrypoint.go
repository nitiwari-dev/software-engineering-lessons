package app

import (
	"fun-with-golang/config"
	"fun-with-golang/intro"
	"github.com/rs/zerolog/log"
	"os"
)

func StartApp() {
	appConfig, err := config.InitAppConfig()
	if err != nil {
		log.Err(err).Msgf("failed to init the app config")
		os.Exit(1)
	}

	//Check to enable intro
	if appConfig.ServiceConfig.IntroConfig.IsEnabled {
		intro.Intro()
	}
}
