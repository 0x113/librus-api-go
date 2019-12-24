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
	User  *User              `json:"User"`
	Class *ResourceReference `json:"Class"`
}

type User struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type ResourceReference struct {
	ID  int    `json:"Id"`
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

type GradesResponse struct {
	Grades *[]Grade
}

type Grade struct {
	Subject               *ResourceReference `json:"Subject"`
	Category              *ResourceReference `json:"Category"`
	AddedBy               *ResourceReference `json:"AddedBy"`
	Grade                 string             `json:"Grade"`
	Date                  string             `json:"Date"`
	AddDate               string             `json:"AddDate"`
	Semester              uint8              `json:"Semester"`
	IsConstituent         bool               `json:"IsConstituent"`
	IsSemester            bool               `json:"IsSemester"`
	IsSemesterProposition bool               `json:"IsSemesterProposition"`
	IsFinal               bool               `json:"IsFinal"`
	IsFinalProposition    bool               `json:"IsFinalProposition"`
}

type GradeDetails struct {
	Subject               *Subject  `json:"Subject"`
	Category 							*Category `json:"Category"`
	Grade                 string    `json:"Grade"`
	Date                  string    `json:"Date"`
	AddDate               string    `json:"AddDate"`
	Semester              uint8     `json:"Semester"`
	IsConstituent         bool      `json:"IsConstituent"`
	IsSemester            bool      `json:"IsSemester"`
	IsSemesterProposition bool      `json:"IsSemesterProposition"`
	IsFinal               bool      `json:"IsFinal"`
	IsFinalProposition    bool      `json:"IsFinalProposition"`
}

type SubjectResponse struct {
	Subject *Subject `json:"Subject"`
}

type Subject struct {
	Name string `json:"Name"`
}

type CategoryResponse struct {
	Category *Category `json:"Category"`
}

type Category struct {
	Name string `json:"Name"`
}
