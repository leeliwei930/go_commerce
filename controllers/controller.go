package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/leeliwei930/go_commerce/repository"
)

type ResponsePayload struct {
	Status    int                          `json:"status"`
	Paginator *repository.PaginatorPayload `json:"paginator,omitempty"`
	Data      interface{}                  `json:"data,omitempty"`
	Errors    interface{}                  `json:"errors,omitempty"`
}

func (r *ResponsePayload) JSON(status int, data interface{}, paginator *repository.PaginatorPayload) ([]byte, error) {
	r.Data = data
	r.Status = status
	if paginator != nil {
		r.Paginator = paginator
	}
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
