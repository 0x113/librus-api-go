package librus_api_go

type LuckyNumberResponse struct {
	LuckyNumber *LuckyNumber           `json:"LuckyNumber"`
	Resources   map[string]interface{} `json:"Resources"`
	Url         string                 `json:"Url"`
}

type LuckyNumber struct {
	LuckyNumber    int    `json:"LuckyNumber"`
	LuckyNumberDay string `json:"LuckyNumberDay"`
}
