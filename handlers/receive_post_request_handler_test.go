package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/mrpiggy97/testingAwsBackend/multiplexer"
	"github.com/rs/zerolog/log"
)

func testRecievePostRequestHandler(testCase *testing.T) {

	//run server
	go multiplexer.Runserver()

	//set test data
	var data map[string]string = make(map[string]string)
	data["data"] = "this is the data sent to server"
	jsonData, _ := json.Marshal(data)
	var buffer *bytes.Buffer = bytes.NewBuffer(jsonData)

	//set request and client
	var client *http.Client = &http.Client{}
	request, requestError := http.NewRequest(
		"POST",
		"http://localhost:8000/api/v1/recieve-post-request",
		buffer,
	)
	if requestError != nil {
		testCase.Error(requestError.Error())
	}

	//send request and make tests
	response, responseError := client.Do(request)
	if responseError != nil {
		testCase.Error(responseError.Error())
	}

	if response.StatusCode != 202 {
		testCase.Error(response.Status)
	} else {
		decodedResponse, _ := io.ReadAll(response.Body)
		log.Info().Msg(string(decodedResponse))
	}
}

func TestRecievePostRequestHandler(testCase *testing.T) {
	testCase.Run("action=test-recieve-post-request-handler", testRecievePostRequestHandler)
}
