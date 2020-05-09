package librus_api_go

import "encoding/json"

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
