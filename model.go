package librus_api_go

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

type LibrusHeader struct {
	Key   string
	Value string
}

type Defaults struct {
	Resources map[string]interface{} `json:"Resources"`
	Url       string                 `json:"Url"`
}

type LuckyNumberResponse struct {
	LuckyNumber *LuckyNumber `json:"LuckyNumber"`
	Defaults
}

type LuckyNumber struct {
	LuckyNumber    int    `json:"LuckyNumber"`
	LuckyNumberDay string `json:"LuckyNumberDay"`
}

type UserInfoResponse struct {
	LibrusMe *LibrusMe `json:"Me"`
	Defaults
}

type LibrusMe struct {
	Account map[string]interface{} `json:"Account"`
	Refresh int										 `json:"Refresh"`
	User 		*User									 `json:"User"`
	Class 	map[string]interface{} `json:"Class"`
}

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
