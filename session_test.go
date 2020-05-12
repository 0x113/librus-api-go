package librus_api_go_test

import (
	"errors"
	"net/http"
	"testing"

	golibrus "github.com/0x113/librus-api-go"
	"github.com/0x113/librus-api-go/mocks"

	"github.com/stretchr/testify/assert"
)

var host = "https://api.librus.pl/"

func TestSuccessGetData(t *testing.T) {
	l := &golibrus.Librus{}
	expectedRes := &http.Response{}
	expectedReq, err := http.NewRequest("GET", host+"2.0/non-existent-endpoint", nil)
	assert.Nil(t, err)
	expectedReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	expectedReq.Header.Set("Authorization", "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=")

	successClient := &mocks.MockClient{
		ExpectedRes: expectedRes,
		Err:         nil,
	}
	l.Client = successClient

	res, err := l.GetData("non-existent-endpoint")

	assert.Equal(t, successClient.Req, expectedReq)
	assert.Equal(t, expectedRes, res)
	assert.Nil(t, err)
}

func TestFailGetData(t *testing.T) {
	l := &golibrus.Librus{}
	expectedRes := &http.Response{}
	expectedReq, err := http.NewRequest("POST", host+"2.0/non-existent-endpoint", nil)
	assert.Nil(t, err)
	expectedReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	expectedReq.Header.Set("Authorization", "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=")

	client := &mocks.MockClient{
		ExpectedRes: expectedRes,
		Err:         errors.New("Wrong method"),
	}
	l.Client = client

	res, err := l.GetData("non-existent-endpoint")

	assert.NotEqual(t, client.Req, expectedReq)
	assert.Equal(t, expectedRes, res)
	assert.NotNil(t, err)
}
