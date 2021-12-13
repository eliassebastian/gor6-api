package util

import (
	"encoding/json"
	"net/http"
	"time"
)

type output struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
	Error     bool        `json:"error"`
	Payload   interface{} `json:"payload"`
}

func ReturnJSON(w http.ResponseWriter, status int, m string, data interface{}) {
	message := output{
		Code:      status,
		Message:   m,
		Timestamp: time.Now(),
		Error:     data == nil,
		Payload:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func ErrorJSON(w http.ResponseWriter, status int, m string) {
	ReturnJSON(w, status, m, nil)
}
