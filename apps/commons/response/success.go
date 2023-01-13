package response

import (
	"net/http"
)

type Success struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Payload    interface{} `json:"payload,omitempty"`
}

var (
	generalSuccess = Success{
		StatusCode: http.StatusOK,
		Message:    "success",
	}
	createSuccess = Success{
		StatusCode: http.StatusCreated,
		Message:    "created success",
	}
)

func GeneralSuccessCustomMessageAndPayload(message string, payload interface{}) *Success {
	succ := generalSuccess
	succ.Message = message
	succ.Payload = payload

	return &succ
}

func CreatedSuccessCustomMessagePayload(message string, payload interface{}) *Success {
	succ := createSuccess
	succ.Message = message
	succ.Payload = payload

	return &succ
}
