package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrpiggy97/testingAwsBackend/loggers"
)

func RecieveDeleteRequest(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loggers.HandlerLogger(request, "RecieveDeleteRequest")
	var data map[string]string = make(map[string]string)
	data["data"] = "delete request accepted"
	jsonData, _ := json.Marshal(data)
	writer.Header().Add("Content-type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(jsonData)
}
