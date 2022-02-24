package multiplexer

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/testingAwsBackend/handlers"
	"github.com/mrpiggy97/testingAwsBackend/middlewares"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Router                 *httprouter.Router
	allowedCrossSiteOrigin string
}

func (serverInstance *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	message := fmt.Sprintf("server recieved request from remote address %v", request.RemoteAddr)
	writer.Header().Add("Access-Control-Allow-Origin", serverInstance.allowedCrossSiteOrigin)
	log.Info().Msg(message)
	request = middlewares.ApplyMiddlewares(request)
	serverInstance.Router.ServeHTTP(writer, request)
}

func NewServer() *Server {
	var multiplexer *Server = &Server{
		Router:                 httprouter.New(),
		allowedCrossSiteOrigin: os.Getenv("ALLOWED_CROSS_SITE_ORIGIN"),
	}
	multiplexer.Router.GET("/api/v1/get-random-number", handlers.RandomNumberHandler)
	multiplexer.Router.POST("/api/v1/recieve-post-request", handlers.RecievePostRequest)
	multiplexer.Router.PUT("/api/v1/recieve-put-request", handlers.RecievePutRequest)
	multiplexer.Router.DELETE("/api/v1/recieve-delete-request", handlers.RecieveDeleteRequest)
	return multiplexer
}
