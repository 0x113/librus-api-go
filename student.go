package librus_api_go

import (
	"encoding/json"
	"strconv"
)

// GetUserInfo calls https://api.librus.pl/2.0/Me and returns info from this endpoint
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
