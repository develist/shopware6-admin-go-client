package shopware6_admin_go_client

import (
	"encoding/json"
	"errors"
	"github.com/develist/shopware6-admin-go-client/entity"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toKebabCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}-${2}")
	return strings.ToLower(snake)
}

func queryArray[T any](httpClient *http.Client, request *http.Request) (*[]T, error) {
	body, err := getBody(httpClient, request)
	if body == nil || err != nil {
		return nil, err
	}

	var results entity.Results[T]
	if err := json.Unmarshal(*body, &results); err != nil {
		return nil, err
	}

	return &results.Data, nil
}

func getBody(httpClient *http.Client, request *http.Request) (*[]byte, error) {
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	resp, err := httpClient.Do(request)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
		return &body, nil
	case 404:
		return nil, nil
	default:
		return nil, errors.New(string(body[:]))
	}
}
