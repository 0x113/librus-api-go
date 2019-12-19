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
	key 	string
	value string
}

var host = "https://api.librus.pl/"

var headers = []librusHeader{
	{
		key: "Authorization",
		value: "Basic Mjg6ODRmZGQzYTg3YjAzZDNlYTZmZmU3NzdiNThiMzMyYjE=",
	},
	{
		key:   "Content-Type",
		value: "application/x-www-form-urlencoded",
	},
}

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
	for _, h := range headers {
		req.Header.Set(h.key, h.value)
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
	headers[0].value = "Bearer " + okResponse.AccessToken

	return nil
}
