package librus_api_go

import (
	"encoding/json"
	"strconv"
)

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
