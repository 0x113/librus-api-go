package librus_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var host = "https://api.librus.pl/"

var Headers = []LibrusHeader{
	{
		Key:   "Authorization",
		Value: "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=",
	},
	{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	},
}

func (l *Librus) CreateSession() error {
	postData := url.Values{}
	postData.Set("username", l.Username)
	postData.Set("password", l.Password)
	postData.Set("librus_long_term_token", "1")
	postData.Set("grant_type", "password")

	// new http client
	client := &http.Client{}

	// request
	req, err := http.NewRequest("POST", host+"OAuth/Token", strings.NewReader(postData.Encode()))
	// add headers
	for _, h := range Headers {
		req.Header.Set(h.Key, h.Value)
	}

	if err != nil {
		return err
	}

	// response
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check response code
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error status code, wanted: %v, got: %v", http.StatusOK, res.StatusCode)
	}

	// decode json response
	okResponse := new(OKResponse)
	err = json.NewDecoder(res.Body).Decode(&okResponse)
	if err != nil {
		return err
	}

	// change authorization header
	Headers[0].Value = "Bearer " + okResponse.AccessToken

	return nil
}

// GetData returns data from url e.g. https://api.librus.pl/2.0/LuckyNumbers
func (l *Librus) GetData(url string) (*http.Response, error) {
	// new http client
	client := &http.Client{}

	// request
	req, err := http.NewRequest("GET", host+"2.0/"+url, nil)
	// add headers
	for _, h := range Headers {
		req.Header.Set(h.Key, h.Value)
	}

	if err != nil {
		return nil, err
	}

	// response
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
