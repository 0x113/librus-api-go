package librus_api_go

import (
	"encoding/json"
	"strconv"
	"sync"
)

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
