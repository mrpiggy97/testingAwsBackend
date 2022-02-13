package loggers

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// HandlerLogger will log basic info about request
func HandlerLogger(request *http.Request, handlerName string) {
	var remoteAddress string = request.RemoteAddr
	var message string = fmt.Sprintf("%v handler called from address %v", handlerName, remoteAddress)
	log.Info().Msg(message)
}
