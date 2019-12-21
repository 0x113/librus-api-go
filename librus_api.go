package librus_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Librus struct {
	Username string
	Password string
}

type OKResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	AccountGroup int    `json:"account_group"`
}

type librusHeader struct {
	Key 	string
	Value string
}

var host = "https://api.librus.pl/"

var Headers = []librusHeader{
	{
		Key: "Authorization",
		Value: "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=",
	},
	{
		Key:   "Content-Type",
		Value: "application/x-www-form-urlencoded",
	},
}

// Login method returns authorization token
func (l *Librus) Login() error {
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

func (l *Librus) GetLuckyNumber() (*LuckyNumber, error) {
	res, err := l.GetData("LuckyNumbers")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get lucky number
	luckyNumber := new(LuckyNumberResponse)
	err = json.NewDecoder(res.Body).Decode(&luckyNumber)
	if err != nil {
		return nil, err
	}

	return luckyNumber.LuckyNumber, nil
}
