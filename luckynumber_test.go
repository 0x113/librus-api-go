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

func TestSuccessGetLuckyNumber(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"LuckyNumber": {"LuckyNumber": 1, "LuckyNumberDay": "2020-01-03"}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	luckyNumber, err := l.GetLuckyNumber()

	assert.Nil(t, err)
	assert.Equal(t, 1, luckyNumber.LuckyNumber)
	assert.Equal(t, "2020-01-03", luckyNumber.LuckyNumberDay)
}

func TestFailGetLuckyNumber(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"LuckyNumber": "1", "LuckyNumberDay": "2029-01-93"}` // invalid json (LuckyNumber is string instead of int)
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	luckyNumber, err := l.GetLuckyNumber()
	assert.NotNil(t, err)
	assert.Nil(t, luckyNumber)
}
