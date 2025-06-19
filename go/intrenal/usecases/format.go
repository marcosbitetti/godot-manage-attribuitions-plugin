package usecases

import (
	"encoding/json"
)

const SuccessMsg = "done"

func FormatJSON(data interface{}, err error) []byte {
	if err != nil {
		return []byte(`{"status":"error", "message":"` + err.Error() + `"}`)
	}
	type Response struct {
		Status  string      `json:"status"`
		Message *string     `json:"message,omitempty"`
		Data    interface{} `json:"data"`
	}
	response := Response{
		Status:  "success",
		Message: nil,
		Data:    data,
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte(`{"status":"error", "message":"` + err.Error() + `"}`)
	}
	return bytes
}
