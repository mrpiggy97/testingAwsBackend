package multiplexer

import (
	"fmt"
	"net/http"

	"github.com/mrpiggy97/testingAwsBackend/handlers"
	"github.com/rs/zerolog/log"
)

func Runserver() {
	var multiplexer *Server = NewServer()
	var address string = "0.0.0.0:8000"
	var message string = fmt.Sprintf("server started listening at address %v", address)
	log.Info().Msg(message)
	multiplexer.Router.GET("/get-random-number", handlers.RandomNumberHandler)
	multiplexer.Router.POST("/recieve-post-request", handlers.RecievePostRequest)
	multiplexer.Router.PUT("/recieve-put-request", handlers.RecievePutRequest)
	multiplexer.Router.DELETE("/recieve-delete-request", handlers.RecieveDeleteRequest)
	http.ListenAndServe(address, multiplexer)
}
