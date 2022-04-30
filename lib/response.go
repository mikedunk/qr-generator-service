package lib

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(w, "%s", err.Error())
	}

}

func NewResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, Response{
		Success: false,
		Error:   err.Error(),
		Data:    nil,
	})
}

func IMAGEBYTES(w http.ResponseWriter, statusCode int, data []byte) {
	w.Header().Set("Content-Type", "octet-stream")
	w.WriteHeader(statusCode)

}
