package api

import "github.com/rs/zerolog"

var log zerolog.Logger

func OnInit(logger zerolog.Logger) {
	log = logger
}
