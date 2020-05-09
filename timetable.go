package librus_api_go

import (
	"encoding/json"
	"sync"
)

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
