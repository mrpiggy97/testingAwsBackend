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
	multiplexer.router.GET("/get-random-number", handlers.RandomNumberHandler)
	multiplexer.router.POST("/recieve-post-request", handlers.RecievePostRequest)
	multiplexer.router.PUT("/recieve-put-request", handlers.RecievePutRequest)
	multiplexer.router.DELETE("/recieve-delete-request", handlers.RecieveDeleteRequest)
	http.ListenAndServe(address, multiplexer)
}
