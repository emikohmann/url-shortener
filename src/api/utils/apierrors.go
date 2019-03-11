package utils

// Use this struct to represent REST API Errors
type ApiError struct {
    Error      string `json:"error"`
    StatusCode int    `json:"status_code"`
}
