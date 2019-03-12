package apierrors

import (
    "strings"
)

const (
    errDuplicatedEntry = "duplicate entry"
)

// Use this struct to represent REST API Errors
type ApiError struct {
    Error      string `json:"error"`
    StatusCode int    `json:"status_code"`
}

func (apiErr *ApiError) IsDuplicatedEntryError() bool {
    return strings.Contains(
        strings.ToLower(apiErr.Error),
        errDuplicatedEntry,
    )
}
