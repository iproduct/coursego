package book

import (
	"encoding/json"
	"net/http"
	"time"
)

type SetResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	AccessTime string      `json:"accessTime"`
}

func HttpResponseSuccess(w http.ResponseWriter, r *http.Request, data interface{}) {
	setResponse := SetResponse{
		Status:     http.StatusText(200),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func HttpResponseError(w http.ResponseWriter, r *http.Request, data interface{}, code int) {
	setResponse := SetResponse{
		Status:     http.StatusText(code),
		AccessTime: time.Now().Format("02-01-2006 15:04:05"),
		Data:       data}
	response, _ := json.Marshal(setResponse)
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(code)
	w.Write(response)
}
