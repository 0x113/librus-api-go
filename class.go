package librus_api_go

import (
	"encoding/json"
	"strconv"
	"sync"
)

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
