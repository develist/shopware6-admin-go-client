package shopware6_admin_go_client

import (
	"encoding/json"
	"github.com/develist/shopware6-admin-go-client/response"
	"github.com/develist/shopware6-admin-go-client/search"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const contentTypeHeader string = "Content-Type"
const acceptHeader string = "Accept"
const mimeTypeJson string = "application/json"

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toKebabCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

func queryArray[T any](httpClient *http.Client, request *http.Request) (*[]T, *response.ApiErrors) {
	body, err := getBody(httpClient, request)
	if err != nil {
		return nil, err
	}
	var results search.Results[T]
	if err := json.Unmarshal(*body, &results); err != nil {
		return nil, createAPIErrors(err)
	}
	return &results.Data, nil
}

func querySingle[T any](httpClient *http.Client, request *http.Request) (*T, *response.ApiErrors) {
	body, err := getBody(httpClient, request)
	if err != nil {
		return nil, err
	}
	var result search.Result[T]
	if err := json.Unmarshal(*body, &result); err != nil {
		return nil, createAPIErrors(err)
	}
	return &result.Data, nil
}

func getBody(httpClient *http.Client, request *http.Request) (*[]byte, *response.ApiErrors) {
	request.Header.Set(contentTypeHeader, mimeTypeJson)
	request.Header.Set(acceptHeader, mimeTypeJson)
	resp, err := httpClient.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, createAPIErrors(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, createAPIErrors(err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		return &body, nil
	case http.StatusNoContent:
		return &body, nil
	default:
		return nil, createAPIErrorsBytes(&body)
	}
}

func createAPIErrorsBytes(err *[]byte) *response.ApiErrors {
	if err == nil {
		return nil
	}
	var result response.ApiErrors
	if unmarshalErr := json.Unmarshal(*err, &result); unmarshalErr != nil {
		result.Raw = string(*err)
		return &result
	} else {
		result.Raw = string(*err)
		return &result
	}
}

func createAPIErrors(err error) *response.ApiErrors {
	data := []byte(err.Error())
	return createAPIErrorsBytes(&data)
}
