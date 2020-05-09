package librus_api_go

import (
	"encoding/json"
	"strconv"
)

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
