package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/testingAwsBackend/loggers"
	"github.com/mrpiggy97/testingAwsBackend/middlewares"
)

func RandomNumberHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loggers.HandlerLogger(request, "RandomNumberHandler")
	var key middlewares.ContextKey = "is_authenticated"
	fmt.Println(request.Context().Value(key))
	var randomNumber int64 = rand.Int63()
	var response map[string]int64 = make(map[string]int64)
	response["randomNumber1"] = randomNumber
	response["randomNumber2"] = randomNumber + 10000
	//response["randomNumber2"] = randomNumber + 100000
	jsonResponse, encodingErr := json.Marshal(response)
	if encodingErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	} else {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusAccepted)
		writer.Write(jsonResponse)
	}
}
