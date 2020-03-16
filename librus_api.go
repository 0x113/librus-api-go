package librus_api_go

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
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
	fmt.Println(okResponse.AccessToken)
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

// GetLuckyNumber returns lucky number for certain day
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

func (l *Librus) GetUserInfo() (*LibrusMe, error) {
	res, err := l.GetData("Me")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get user info
	userInfo := new(UserInfoResponse)
	err = json.NewDecoder(res.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo.LibrusMe, nil
}

// GetUserRealName returns real first and last name of user
func (l *Librus) GetUserRealName() (string, error) {
	userInfo, err := l.GetUserInfo()
	if err != nil {
		return "", err
	}

	return userInfo.User.FirstName + " " + userInfo.User.LastName, nil
}

// GetUserClass returns details of the class to which user belongs
func (l *Librus) GetUserClass() (*ClassDetails, error) {
	userInfo, err := l.GetUserInfo()
	if err != nil {
		return nil, err
	}

	// get class details response
	res, err := l.GetData("Classes/" + strconv.Itoa(userInfo.Class.ID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get class details
	classDetails := new(ClassDetailsResponse)
	err = json.NewDecoder(res.Body).Decode(&classDetails)
	if err != nil {
		return nil, err
	}

	return classDetails.ClassDetails, nil
}

// GetUserGrades returns user grades
func (l *Librus) GetUserGrades() ([]*GradeDetails, error) {
	res, err := l.GetData("Grades")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// grades response
	gradesResponse := new(GradesResponse)
	err = json.NewDecoder(res.Body).Decode(gradesResponse)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	// detailed grades slice
	var detailedGrades []*GradeDetails

	// get grade subject
	for _, g := range *gradesResponse.Grades {
		grade := new(GradeDetails)
		wg.Add(1)

		// set default grades values
		grade.GradeDefault = g.GradeDefault

		// get subject
		go func(subjectID, categoryID, addedByID int, grade *GradeDetails) {
			defer wg.Done()

			// SUBJECT
			subject, err := l.GetSubject(subjectID)
			if err != nil {
				return
			}
			grade.Subject = subject // set grade subject

			// CATEGORY
			category, err := l.GetGradeCategory(categoryID)
			if err != nil {
				return
			}
			grade.Category = category // set grade category

			// ADDED BY
			addedBy, err := l.GetUser(addedByID)
			if err != nil {
				return
			}
			grade.AddedBy = addedBy

			// append grade
			detailedGrades = append(detailedGrades, grade)
		}(g.Subject.ID, g.Category.ID, g.AddedBy.ID, grade)

	}
	wg.Wait()

	return detailedGrades, nil
}

func (l *Librus) GetSubject(id int) (*Subject, error) {
	res, err := l.GetData("Subjects/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get subject data
	subjectResponse := new(SubjectResponse)
	err = json.NewDecoder(res.Body).Decode(subjectResponse)
	if err != nil {
		return nil, err
	}

	return subjectResponse.Subject, nil
}

func (l *Librus) GetGradeCategory(id int) (*Category, error) {
	res, err := l.GetData("Grades/Categories/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get category data
	categoryResponse := new(CategoryResponse)
	err = json.NewDecoder(res.Body).Decode(categoryResponse)
	if err != nil {
		return nil, err
	}

	return categoryResponse.Category, nil
}

func (l *Librus) GetUser(id int) (*User, error) {
	res, err := l.GetData("Users/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get AddedBy response
	addedByResponse := new(AddedByResponse)
	err = json.NewDecoder(res.Body).Decode(addedByResponse)
	if err != nil {
		return nil, err
	}

	return addedByResponse.AddedBy, nil
}

func (l *Librus) GetAttendance() (error, error) {
	// Types
	// 1   - nieobocność
	// 2   - spóźnienie
	// 3   - nieobocność usp.
	// 4   - zwolnienie
	// 100 - obecność
	return nil, nil
}
