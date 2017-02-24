package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PostHTTP converts the struct to an HTTP request, sends the request, and reads the response
func PostHTTP(ipAddress string, payload Body) ([]byte, error) {

	log.Printf("Sending Request")

	// create json payload
	postBody, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, err
	}

	// create address
	address := fmt.Sprintf("http://%s/cgi-bin/web.cgi", ipAddress)

	//create request
	request, err := http.NewRequest("POST", address, bytes.NewBuffer(postBody))
	if err != nil {
		return []byte{}, err
	}

	// set headers
	request.Header.Set("Content-Type", "application/json")

	// do request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	//get response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	} else if responseBody == nil {
		return []byte{}, errors.New("Blank response body")
	}

	return responseBody, nil
}

func BuildAndSendPayload(ipAddress string, requestType string, requestCategory string, parameterCategory string, parameterCode string) error {

	log.Printf("Building JSON")

	request := Request{
		Type:     requestType,
		Category: requestCategory,
	}

	parameter := Parameter{
		Category: parameterCategory,
		Code:     parameterCode,
	}
	payload := Body{
		Req:   request,
		Param: parameter,
	}

	_, err := PostHTTP(ipAddress, payload)
	if err != nil {
		return err
	}

	return nil
}
