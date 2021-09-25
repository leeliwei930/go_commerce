package controllers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
}

type ResponsePayload struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func (r *ResponsePayload) JSON(status int, data interface{}) ([]byte, error) {
	r.Data = data
	r.Status = status
	return json.Marshal(r)
}

func (r *ResponsePayload) Error(status int, e error) ([]byte, error) {
	r.Errors = e.Error()
	r.Status = status
	return json.Marshal(r)
}

func Response() *ResponsePayload {
	return &ResponsePayload{}
}

func InjectJSONResponseHeader(next http.Handler) http.Handler {

	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(response, request)
	})
}
