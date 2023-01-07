package transport

import "fmt"

// ErrorResponse represents error object to keep server error.
type ErrorResponse struct {
	Code    int    `json:"-"`
	Message string `json:"error_message"`
}

// Error implements Error interface.
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("api returned an error with status code: %d message %s", e.Code, e.Message)
}

// GetCode returns original Status Code.
func (e ErrorResponse) GetCode() int {
	return e.Code
}

// GetMessage returns original error message
func (e ErrorResponse) GetMessage() string {
	return e.Message
}
