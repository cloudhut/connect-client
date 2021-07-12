package connect

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	hostUrl := "https://kafka-connect.com"
	timeOut := 5 * time.Second
	c := NewClient(WithHost(hostUrl), WithTimeout(timeOut))

	assert.Equal(t, hostUrl, c.hostURL)
	assert.Equal(t, timeOut, c.timeout)
}

// newJsonStringResponder creates a Responder from a given body (as a string) and status code and
// sets the JSON content type.
//
// To pass the content of an existing file as body use httpmock.File as in:
//   httpmock.NewStringResponder(200, httpmock.File("body.txt").String())
func newJsonStringResponder(status int, body string) httpmock.Responder {
	return httpmock.ResponderFromResponse(newJsonStringResponse(status, body))
}

func newJsonStringResponse(status int, body string) *http.Response {
	h := http.Header{}
	h.Set("Content-Type", "application/json")
	return &http.Response{
		Status:        strconv.Itoa(status),
		StatusCode:    status,
		Body:          httpmock.NewRespBodyFromString(body),
		Header:        h,
		ContentLength: -1,
	}
}
