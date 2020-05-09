package librus_api_go

import (
	"encoding/json"
	"sync"
)

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
