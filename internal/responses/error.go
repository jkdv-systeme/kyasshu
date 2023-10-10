package responses

import (
	"strconv"
)

func NewError(status int, message string) error {
	return &ServerError{
		Status:  status,
		Message: message,
	}
}

func NewValidationError(status int, message string, fields []FieldError) error {
	return &ValidationError{
		ServerError: ServerError{
			Status:  status,
			Message: message,
		},
		Fields: fields,
	}
}

type ServerError struct {
	Status  int
	Message string
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationError struct {
	ServerError
	Fields []FieldError
}

func (e *ServerError) Error() string {
	return strconv.Itoa(e.Status) + " - " + e.Message
}

type ErrorResponse struct {
	Status    int          `json:"status"`
	Message   string       `json:"message"`
	RequestID string       `json:"request_id,omitempty"`
	Fields    []FieldError `json:"fields,omitempty"`
} //@name ErrorResponse
