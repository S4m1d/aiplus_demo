package main

import (
	"aiplus_demo/src/api"
	"aiplus_demo/src/bl"
	"aiplus_demo/src/da"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	configureGlobalLogger()
	/*
	*	by passing logger in each module instead of using global logger everywhere we keep ability to
	*	provide app with more sophisticated logging in future without hard refactoring,
	*	for example if different modules could require to have separate outputs or levels
	 */
	da.OnInit(log.Logger)
	defer da.OnClose()
	bl.OnInit(log.Logger)
	api.OnInit(log.Logger)

	api.ConfigureRoutes()
}

func configureGlobalLogger() {
	//todo parametrize
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
