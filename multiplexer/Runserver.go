package multiplexer

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func Runserver() {
	var multiplexer *Server = NewServer()
	var address string = "0.0.0.0:8000"
	var message string = fmt.Sprintf("server started listening at address %v", address)
	log.Info().Msg(message)
	http.ListenAndServe(address, multiplexer)
}
