package mocks

import (
	"net/http"
)

type MockClient struct {
	ExpectedRes *http.Response
	Req         *http.Request
	Err         error
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.Req = req
	return m.ExpectedRes, m.Err
}
