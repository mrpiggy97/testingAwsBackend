package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/testingAwsBackend/loggers"
	"github.com/rs/zerolog/log"
)

func RecievePostRequest(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//log request data sent
	requestData, _ := io.ReadAll(request.Body)
	log.Info().Msg(string(requestData))
	loggers.HandlerLogger(request, "RecievePostRequest")
	var data map[string]string = make(map[string]string)
	data["data"] = "request accepted"
	jsonData, _ := json.Marshal(data)
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(jsonData)
}
