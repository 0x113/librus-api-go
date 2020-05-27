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

func TestSuccessGetLesson(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `{"Lesson": {"Teacher": {"ID": 1, "Url": "/teacher-endpoint"}, "Subject": {"ID": 2, "Url": "/subject-endpoint"}, "Class": {"ID": 3, "Url": "/class-endpoint"}}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	expectedLesson := &golibrus.Lesson{
		Teacher: &golibrus.ResourceReference{ID: 1, Url: "/teacher-endpoint"},
		Subject: &golibrus.ResourceReference{ID: 2, Url: "/subject-endpoint"},
		Class:   &golibrus.ResourceReference{ID: 3, Url: "/class-endpoint"},
	}

	lesson, err := l.GetLesson(123)
	assert.Nil(t, err)
	assert.Equal(t, expectedLesson, lesson)
}

func TestFailGetLesson(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			json := `` // invalid json
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(json))),
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	lesson, err := l.GetLesson(123)
	assert.NotNil(t, err)
	assert.Nil(t, lesson)
}

func TestInternalServerErrorGetLesson(t *testing.T) {
	client := &mocks.MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusInternalServerError,
			}, nil
		},
	}
	l := &golibrus.Librus{Client: client}

	// set authorization token
	golibrus.Headers[0].Value = "Bearer HESOYAM"

	lesson, err := l.GetLesson(123)
	assert.NotNil(t, err)
	assert.Nil(t, lesson)
}
