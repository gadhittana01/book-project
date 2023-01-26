package resthttp

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type baseResp struct {
	startTime   time.Time
	Data        interface{} `json:"data"`
	ElapsedTime string      `json:"elapsed_time"`
	RequestID   string      `json:"request_id"`
	IsError     bool        `json:"is_error"`
}

const (
	InvalidRequestParam = "Invalid Request Parameter"
	IncompleteParam     = "Parameter tidak lengkap"
	DataNotFound        = "Data tidak ditemukan"
)

func newResponse(startTime time.Time) baseResp {
	return baseResp{
		startTime: startTime,
	}
}

func (br *baseResp) setElapsedTime() {
	br.ElapsedTime = time.Since(br.startTime).String()
}

func (br *baseResp) setBadRequest(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Bad Request"
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setBadRequest error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(respBytes)
}

func (br *baseResp) setInternalServerError(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Internal server error"
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
		"status":        400,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setInternalServerError error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respBytes)
}

func (br *baseResp) setOK(data interface{}, w http.ResponseWriter) {
	br.Data = data
	br.setElapsedTime()

	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setOK error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}

func (br *baseResp) setNotFound(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Not found"
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setNotFound error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(respBytes)
}

func (br *baseResp) setForbidden(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Forbidden."
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, _ := json.Marshal(br)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write(respBytes)
}

func (br *baseResp) setBadRequestWithStatus(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Bad Request"
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
		"status":        http.StatusBadRequest,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setBadRequest error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(respBytes)
}

func (br *baseResp) setInternalServerErrorWithStatus(msg string, w http.ResponseWriter) {
	if msg == "" {
		msg = "Internal server error"
	}
	br.Data = map[string]interface{}{
		"error_message": msg,
		"status":        http.StatusInternalServerError,
	}
	br.setElapsedTime()
	br.IsError = true
	respBytes, err := json.Marshal(br)
	if err != nil {
		log.Println(br.RequestID, "setInternalServerError error : %+v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respBytes)
}
