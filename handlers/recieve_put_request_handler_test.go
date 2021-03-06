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

func testRecievePutRequestHandler(testCase *testing.T) {

	//set data to send
	var data map[string]string = make(map[string]string)
	data["data"] = "this is a test"
	jsonData, _ := json.Marshal(data)
	var buffer *bytes.Buffer = bytes.NewBuffer(jsonData)

	//run server
	go multiplexer.Runserver()

	//set request and client
	var client *http.Client = &http.Client{}
	request, requestError := http.NewRequest(
		"PUT",
		"http://localhost:8000/api/v1/recieve-put-request",
		buffer,
	)

	if requestError != nil {
		testCase.Error(requestError.Error())
	}

	//make request and test
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

func TestRecievePutRequestHandler(testCase *testing.T) {
	testCase.Run("action=test-recieve-put-request-handler", testRecievePutRequestHandler)
}
