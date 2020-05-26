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

func TestSuccessGetSchoolInfo(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"School": {"Name": "Cool school", "Town": "Cracow", "Street": "Sezamkowa", "State": "-", "BuildingNumber": "12"}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	expectedSchool := &golibrus.School{
		Name:           "Cool school",
		Town:           "Cracow",
		Street:         "Sezamkowa",
		State:          "-",
		BuildingNumber: "12",
	}

	schoolInfo, err := l.GetSchoolInfo()
	assert.Nil(t, err)
	assert.Equal(t, expectedSchool, schoolInfo)
}

func TestFailGetSchoolInfo(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"School": {"Name": "Cool school", "Town": "Cracow", "Street": "Sezamkowa", "State": "-", "BuildingNumber": 12}}` // invalid json (BuildingNumber should be string not int)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	schoolInfo, err := l.GetSchoolInfo()
	assert.NotNil(t, err)
	assert.Nil(t, schoolInfo)
}
