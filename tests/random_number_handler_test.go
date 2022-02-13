package tests

import (
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/mrpiggy97/testingAwsBackend/httpReader"
	"github.com/mrpiggy97/testingAwsBackend/multiplexer"
)

func TestRandomNumberHandler(testCase *testing.T) {
	go multiplexer.Runserver()
	//slee to give server time to start
	time.Sleep(time.Second * 1)
	res, resErr := http.Get("http://localhost:8000/get-random-number")
	decodedResBody, _ := io.ReadAll(res.Body)
	var mapResBody map[string]string = httpReader.BodyReader(decodedResBody)
	if resErr != nil {
		testCase.Error(resErr)
	}

	firstNumber, atoiErr := strconv.Atoi(mapResBody["randomNumber1"])
	if atoiErr != nil {
		testCase.Error(atoiErr)
	}
	secondNumber, _ := strconv.Atoi(mapResBody["randomNumber2"])
	if firstNumber+10000 != secondNumber {
		testCase.Error(
			"response from server not the expected, randomNumber2 should be 10000 more than randomNumber1",
		)
	}
}
