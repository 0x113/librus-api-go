package librus_api_go_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	golibrus "github.com/0x113/librus-api-go"
	"github.com/0x113/librus-api-go/mocks"

	"github.com/stretchr/testify/assert"
)

var host = "https://api.librus.pl/"

func TestSuccessCreateSession(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"access_token": "HESOYAM","token_type": "basic", "expires_in": 60, "refresh_token": "motherlode", "account_group": 1}`

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}

	expectedHeaders := []golibrus.LibrusHeader{
		{
			Key:   "Authorization",
			Value: "Bearer HESOYAM",
		},
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	}

	l := &golibrus.Librus{
		Client: client,
	}

	err := l.CreateSession() // this will update global variable named 'Headers'
	assert.Equal(t, expectedHeaders, golibrus.Headers)
	assert.Nil(t, err)
}

func TestFailCreateSession(t *testing.T) {
	// reset headers
	golibrus.Headers = []golibrus.LibrusHeader{
		{
			Key:   "Authorization",
			Value: "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=",
		},
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	}

	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(nil),
			}, nil
		},
	}

	l := &golibrus.Librus{
		Client: client,
	}

	expectedHeaders := []golibrus.LibrusHeader{
		{
			Key:   "Authorization",
			Value: "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=",
		},
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	}

	err := l.CreateSession()

	assert.NotNil(t, err)
	assert.Equal(t, expectedHeaders, golibrus.Headers)
}

func TestSuccessGetData(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		},
	}

	expectedHeaders := []golibrus.LibrusHeader{
		{
			Key:   "Authorization",
			Value: "Bearer HESOYAM",
		},
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
	}

	// set authorization token, normally set via CreateSession method
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	l := &golibrus.Librus{Client: client}
	res, err := l.GetData("non-existent-endpoint")
	assert.Equal(t, expectedHeaders, golibrus.Headers)
	assert.Equal(t, 200, res.StatusCode)
	assert.Nil(t, err)
}

func TestFailGetData(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{}, nil
		},
	}

	// change authorization token not to be Bearer
	golibrus.Headers[0].Value = "Basic not-hesoyam"

	l := &golibrus.Librus{Client: client}
	_, err := l.GetData("non-existent-endpoint")
	assert.NotNil(t, err)
}
