package middlewares

import "net/http"

type middlewareFunc func(request *http.Request) *http.Request

var MIDDLEWARES []middlewareFunc = []middlewareFunc{AuthenticationMiddleware}

func ApplyMiddlewares(request *http.Request) *http.Request {
	for _, middleware := range MIDDLEWARES {
		request = middleware((request))
	}
	return request
}
