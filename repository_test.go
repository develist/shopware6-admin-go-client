package shopware6_admin_go_client

import (
	"bytes"
	"github.com/develist/shopware6-admin-go-client/entity"
	"io/ioutil"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestGetDetailErrorResponse(t *testing.T) {
	orderId := "e7ea95d15cf94ce394c4528996efbd55"
	baseUrl := "http://example.com/api"
	errorStr := "ERROR"

	client := NewTestClient(func(req *http.Request) *http.Response {
		// Test request parameters
		wantedUrl := baseUrl + "/order/" + orderId
		if req.URL.String() != wantedUrl {
			t.Fatalf(`Wanted was %v but return value is %v`, wantedUrl, req.URL.String())
		}
		return &http.Response{
			StatusCode: 500,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(errorStr)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})

	// Use Client & URL from our local test server
	repository := CreateEntityRepository[entity.Order](baseUrl, client)
	_, error := repository.GetDetail(orderId)
	if error.Raw != errorStr {
		t.Fatalf(`Wanted was %v but return value is %v`, errorStr, error.Raw)
	}

}
