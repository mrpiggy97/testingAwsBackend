package middlewares

import (
	"context"
	"math/rand"
	"net/http"
)

// a fake authenticator to see how a middleware would be
// implemented to handle authentication
func AuthenticationMiddleware(request *http.Request) *http.Request {
	var randomNumber int64 = rand.Int63()
	var key ContextKey = "is_authenticated"
	var newRequest *http.Request = new(http.Request)
	if randomNumber > 10000000 {
		var newContext context.Context = context.WithValue(
			request.Context(),
			key,
			true,
		)
		newRequest = request.Clone(newContext)
	} else {
		var newContext context.Context = context.WithValue(
			request.Context(),
			key,
			false,
		)
		newRequest = request.Clone(newContext)
	}
	return newRequest
}
