package librus_api_go

import "encoding/json"

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
