package response

import (
	"encoding/json"
	"net/http"
	"time"
)

func SuccessJSON(w http.ResponseWriter, start time.Time, data interface{}) {
	type output struct {
		Status   string        `json:"status"`
		Duration time.Duration `json:"duration"`
		Data     interface{}   `json:"data"`
	}

	message := output{
		Status:   "success",
		Duration: time.Now().Sub(start),
		Data:     data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(message)
}

func FailureJSON(w http.ResponseWriter, m string) {
	type output struct {
		Status string `json:"status"`
		Data   string `json:"data"`
	}

	message := output{
		Status: "failure",
		Data:   m,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(message)
}

func ErrorJSON(w http.ResponseWriter, e error) {
	type output struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	message := output{
		Status:  "error",
		Message: e.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(message)
}
