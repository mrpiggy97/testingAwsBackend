package multiplexer

import (
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/testingAwsBackend/middlewares"
)

type Server struct {
	router                 *httprouter.Router
	allowedCrossSiteOrigin string
	allowedMethods         []string
}

func (serverInstance *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var allowedMethods string = strings.Join(serverInstance.allowedMethods, ",")
	writer.Header().Add("Access-Control-Allow-Origin", serverInstance.allowedCrossSiteOrigin)
	writer.Header().Add("Access-Control-Allow-Methods", allowedMethods)
	request = middlewares.ApplyMiddlewares(request)
	serverInstance.router.ServeHTTP(writer, request)
}

func NewServer() *Server {
	var multiplexer *Server = &Server{
		router:                 httprouter.New(),
		allowedCrossSiteOrigin: os.Getenv("ALLOWED_CROSS_SITE_ORIGIN"),
		allowedMethods:         []string{"GET"},
	}
	return multiplexer
}
