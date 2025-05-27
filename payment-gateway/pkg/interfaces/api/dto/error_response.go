package dto

type ErrorResponse struct {
	Code    string `json:"code"`    // Application-specific error code
	Message string `json:"message"` // User-friendly error message
}

type ValidationErrorResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"` // Field-specific validation errors
}
