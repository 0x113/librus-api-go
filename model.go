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

type LuckyNumberResponse struct {
	LuckyNumber *LuckyNumber `json:"LuckyNumber"`
}

type LuckyNumber struct {
	LuckyNumber    int    `json:"LuckyNumber"`
	LuckyNumberDay string `json:"LuckyNumberDay"`
}

type UserInfoResponse struct {
	LibrusMe *LibrusMe `json:"Me"`
}

type LibrusMe struct {
	User  *User  `json:"User"`
	Class *Class `json:"Class"`
}

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type Class struct {
	ID  int		 `json:"Id"`
	Url string `json:"Url"`
}

type ClassDetailsResponse struct {
	ClassDetails *ClassDetails `json:"Class"`
}

type ClassDetails struct {
	Number           int    `json:"Number"`
	Symbol           string `json:"Symbol"`
	BeginSchoolYear  string `json:"BeginSchoolYear"`
	EndFirstSemester string `json:"EndFirstSemester"`
	EndSchoolYear    string `json:"EndSchoolYear"`
}
