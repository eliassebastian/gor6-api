package util

import (
	"encoding/json"
	"net/http"
	"time"
)

func ReturnJSON(w http.ResponseWriter, start time.Time, status int, m string, data interface{}) {
	type output struct {
		Code      int           `json:"code"`
		Message   string        `json:"message"`
		Timestamp time.Time     `json:"timestamp"`
		Duration  time.Duration `json:"duration"`
		Error     bool          `json:"error"`
		Payload   interface{}   `json:"payload"`
	}

	message := output{
		Code:      status,
		Message:   m,
		Timestamp: time.Now(),
		Duration:  time.Now().Sub(start),
		Error:     data == nil,
		Payload:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}

func ErrorJSON(w http.ResponseWriter, start time.Time, status int, m string) {
	ReturnJSON(w, start, status, m, nil)
}
