package librus_api_go

import (
	"encoding/json"
	"strconv"
)

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
