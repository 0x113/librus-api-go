package mocks

import (
	"net/http"
)

type MockClient struct {
	ExpectedRes *http.Response
	Req         *http.Request
	Err         error
	DoFunc      func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return m.ExpectedRes, m.Err
}
