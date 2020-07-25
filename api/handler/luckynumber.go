package handler

import (
	"encoding/json"
	"net/http"

	client "github.com/0x113/librus-api-go/api/librusclient"
)

// LuckyNumber returns lucky number for certain date in json format
func LuckyNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]interface{})

	luckynumber, err := client.LibrusClient.GetLuckyNumber()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res["error"] = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(luckynumber)
}
