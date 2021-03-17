package cloudconfigclient_test

import (
	"bytes"
	"github.com/Piszmog/cloudconfigclient"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
)

// mockCloudClient is the mocked object that implements CloudClient
type mockCloudClient struct {
	mock.Mock
}

func (c *mockCloudClient) Get(uriVariables ...string) (resp *http.Response, err error) {
	args := c.Called(uriVariables)
	return args.Get(0).(*http.Response), args.Error(1)
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewMockHttpClient creates a mocked HTTP client
func NewMockHttpClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

// NewMockHttpResponse creates a mocked HTTP response
func NewMockHttpResponse(code int, body string) *http.Response {
	return &http.Response{
		StatusCode: code,
		// Send response to be tested
		Body: ioutil.NopCloser(bytes.NewBufferString(body)),
		// Must be set to non-nil value or it panics
		Header: make(http.Header),
	}
}

// NewConfigClient creates a new config client
func NewConfigClient(clients ...cloudconfigclient.CloudClient) cloudconfigclient.ConfigClient {
	return cloudconfigclient.ConfigClient{Clients: clients}
}
