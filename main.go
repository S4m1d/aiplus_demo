package main

import (
	"aiplus_demo/src/api"
	"aiplus_demo/src/bl"
	"aiplus_demo/src/da"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const appLauchErrMsg = "app launch failed"

func main() {
	//loading env variables from .env file
	godotenv.Load()
	configureGlobalLogger()
	/*
	*	by passing logger in each module instead of using global logger everywhere we keep ability to
	*	provide app with more sophisticated logging in future without hard refactoring,
	*	for example if different modules could require to have separate outputs or levels
	 */
	err := da.OnInit(log.Logger)
	defer da.OnClose()
	if err != nil {
		log.Panic().Err(err).Msg(appLauchErrMsg)
	}
	log.Info().Msg("da module initialized")
	bl.OnInit(log.Logger)
	log.Info().Msg("bl module initialized")
	api.OnInit(log.Logger)
	log.Info().Msg("api module initialized")

	err = api.ConfigureRoutesAndListen()
	if err != nil {
		log.Panic().Err(err).Msg(appLauchErrMsg)
	}
}

func configureGlobalLogger() {
	switch os.Getenv(loggerLevelEnv) {
	case logLvlError:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case logLvlWarning:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case logLvlInfo:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case logLvlDebug:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
