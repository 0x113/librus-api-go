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

func TestSuccessGetUser(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"User": {"FirstName": "John", "LastName": "Doe"}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	expectedUser := &golibrus.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	user, err := l.GetUser(123)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestFailGetUser(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{}` // invalid json
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	user, err := l.GetUser(123)
	assert.Nil(t, err)
	assert.Nil(t, user)
}
