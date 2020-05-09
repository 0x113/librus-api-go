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
	err = json.NewDecoder(res.Body).Decode(luckyNumber)
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
	err = json.NewDecoder(res.Body).Decode(userInfo)
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
	err = json.NewDecoder(res.Body).Decode(classDetails)
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

// GetSubject name
func (l *Librus) GetSubject(id int) (string, error) {
	res, err := l.GetData("Subjects/" + strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// get subject data
	subjectResponse := new(SubjectResponse)
	err = json.NewDecoder(res.Body).Decode(subjectResponse)
	if err != nil {
		return "", err
	}

	return subjectResponse.Subject.Name, nil
}

// GetGradeCategory name
func (l *Librus) GetGradeCategory(id int) (string, error) {
	res, err := l.GetData("Grades/Categories/" + strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// get category data
	categoryResponse := new(CategoryResponse)
	err = json.NewDecoder(res.Body).Decode(categoryResponse)
	if err != nil {
		return "", err
	}

	return categoryResponse.Category.Name, nil
}

// GetUser info like first name and last name
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

// GetLesson info like teacher, subject and class
func (l *Librus) GetLesson(id int) (*Lesson, error) {
	res, err := l.GetData("Lessons/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get lesson response
	lessonResponse := new(LessonResponse)
	if err := json.NewDecoder(res.Body).Decode(lessonResponse); err != nil {
		return nil, err
	}

	return lessonResponse.Lesson, nil
}

// GetAttendance retruns list of attendance data like date, lesson number, semester etc
func (l *Librus) GetAttendance() ([]*AttendanceDetails, error) {
	// Types
	// 1   - nieobecność
	// 2   - spóźnienie
	// 3   - nieobecność usp.
	// 4   - zwolnienie
	// 100 - obecność
	res, err := l.GetData("Attendances")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get attendances response
	attendancesResponse := new(AttendanceResponse)
	if err := json.NewDecoder(res.Body).Decode(attendancesResponse); err != nil {
		return nil, err
	}

	// Attendance details
	var wg sync.WaitGroup
	var detailedAttendance []*AttendanceDetails
	for _, a := range attendancesResponse.Attendances {
		attendanceDetails := &AttendanceDetails{}
		// attendance type
		switch a.Type.ID {
		case 1:
			attendanceDetails.Type = "nieobecność"
		case 2:
			attendanceDetails.Type = "spóźnienie"
		case 3:
			attendanceDetails.Type = "nieobecność usp."
		case 4:
			attendanceDetails.Type = "zwolnienie"
		case 100:
			attendanceDetails.Type = "obecność"
		}
		// default fields
		attendanceDetails.Attendance = a.Attendance

		wg.Add(1)
		go func(lessonID, userID int) {
			defer wg.Done()
			// ADDED BY
			teacher, err := l.GetUser(userID)
			if err != nil {
				return
			}
			attendanceDetails.AddedBy = teacher

			// GET SUBJECT
			lesson, err := l.GetLesson(lessonID)
			if err != nil {
				return
			}
			subject, err := l.GetSubject(lesson.Subject.ID)
			if err != nil {
				return
			}
			attendanceDetails.Subject = subject

			detailedAttendance = append(detailedAttendance, attendanceDetails)

		}(a.Lesson.ID, a.AddedBy.ID)
	}
	wg.Wait()

	return detailedAttendance, nil
}

// GetSchoolInfo like name, town, street etc.
func (l *Librus) GetSchoolInfo() (*School, error) {
	res, err := l.GetData("Schools")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// get school response
	schoolResponse := new(SchoolResponse)
	if err := json.NewDecoder(res.Body).Decode(schoolResponse); err != nil {
		return nil, err
	}

	return schoolResponse.School, nil
}

// GetClassFreeDaysTypes returns list of all reasons why class had a free day
func (l *Librus) GetClassFreeDaysTypes() ([]*Type, error) {
	res, err := l.GetData("ClassFreeDays/Types")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	typesResponse := new(TypesResponse)
	if err := json.NewDecoder(res.Body).Decode(typesResponse); err != nil {
		return nil, err
	}

	return typesResponse.Types, nil
}

// GetClassFreeDayTypeById returns reason why class had a free day
func (l *Librus) GetClassFreeDayTypeById(id int) (*Type, error) {
	res, err := l.GetData("ClassFreeDays/Types/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	response := new(TypesResponse)
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return nil, err
	}

	return response.Types[0], nil
}

// GetClassFreeDays returns list of all class free days
func (l *Librus) GetClassFreeDays() ([]*ClassFreeDay, error) {
	res, err := l.GetData("ClassFreeDays")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	classFreeDaysResponse := new(ClassFreeDaysResponse)
	if err := json.NewDecoder(res.Body).Decode(classFreeDaysResponse); err != nil {
		return nil, err
	}

	// get types
	var wg sync.WaitGroup
	freeDays := []*ClassFreeDay{}
	for _, d := range classFreeDaysResponse.ClassFreeDays {
		wg.Add(1)
		freeDay := &ClassFreeDay{}
		freeDay.ClassFreeDayDefault = d.ClassFreeDayDefault
		go func(dayTypeID int) {
			defer wg.Done()
			freeDayType, err := l.GetClassFreeDayTypeById(dayTypeID)
			if err != nil {
				return
			}
			freeDay.Type = freeDayType
			freeDays = append(freeDays, freeDay)
		}(d.Type.ID)
	}
	wg.Wait()

	return freeDays, nil
}

// GetTimetableEntries returns all timetable entries since beginning of school year
func (l *Librus) GetTimetableEntries() ([]*TimetableEntry, error) {
	res, err := l.GetData("TimetableEntries")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	timetableEntriesResponse := new(TimetableEntriesResponse)
	if err := json.NewDecoder(res.Body).Decode(timetableEntriesResponse); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var entries []*TimetableEntry
	for _, te := range timetableEntriesResponse.Entries {
		// get lesson
		wg.Add(1)
		entry := &TimetableEntry{}
		entry.TimetableEntryDefault = te.TimetableEntryDefault
		go func(lessonID int) {
			defer wg.Done()
			lesson, err := l.GetLesson(lessonID)
			if err != nil {
				return
			}
			// get subject
			subject, err := l.GetSubject(lesson.Subject.ID)
			if err != nil {
				return
			}
			entry.Subject = subject

			// get teacher
			teacher, err := l.GetUser(lesson.Teacher.ID)
			entry.Teacher = teacher
			entries = append(entries, entry)

		}(te.Lesson.ID)
	}
	wg.Wait()

	return entries, nil
}
