package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go-erp/domain/throwable"
)

type ErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Error      error  `json:"-"`
}

type SuccessResponse struct {
	StatusCode int         `json:"code"`
	Data       interface{} `json:"data"`
}

type Handler struct {
}

func (h *Handler) Debug(format string, args ...interface{}) {
	if debug := os.Getenv("DEBUG"); debug == "true" {
		log.Printf("[DEBUG] "+format+"\n", args...)
	}
}

func (h *Handler) Error(writer http.ResponseWriter, err error) {
	var statusCode int
	var message string

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch e := err.(type) {
	case *throwable.NotFound, throwable.NotFound:
		statusCode = http.StatusNotFound
		message = e.Error()
	case *throwable.Unauthorized, throwable.Unauthorized:
		statusCode = http.StatusUnauthorized
		message = e.Error()
	case *throwable.InvalidFilter, throwable.InvalidFilter:
		statusCode = http.StatusBadRequest
		message = e.Error()
	case *json.UnsupportedTypeError, *json.UnmarshalTypeError, *json.SyntaxError:
		statusCode = http.StatusBadRequest
		message = "Request body is invalid"
	default:
		statusCode = http.StatusInternalServerError
		message = e.Error()
	}

	h.Debug("%s", err.Error())

	errorResponse := &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	}

	h.Respond(writer, statusCode, errorResponse)
}

func (h *Handler) Respond(writer http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if body, err = json.Marshal(src); err != nil {
		errorBody := "{\"status\": 500, \"message\": \"Something happened wrong during generating response\"}"
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(errorBody))
		return
	}

	writer.WriteHeader(code)
	writer.Write(body)
}
