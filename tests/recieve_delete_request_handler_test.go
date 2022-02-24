package tests

import (
	"io"
	"net/http"
	"testing"

	"github.com/mrpiggy97/testingAwsBackend/multiplexer"
	"github.com/rs/zerolog/log"
)

func TestRecieveDeleteRequestHandler(testCase *testing.T) {

	//run server
	go multiplexer.Runserver()

	//set request and client
	var client *http.Client = &http.Client{}
	request, requestError := http.NewRequest(
		"DELETE",
		"http://localhost:8000/api/v1/recieve-delete-request",
		nil,
	)

	if requestError != nil {
		testCase.Error(requestError.Error())
	}

	//make request and tests
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
