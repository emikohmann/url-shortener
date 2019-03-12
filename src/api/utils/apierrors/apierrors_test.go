package apierrors

import (
    "testing"
    "net/http"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "duplicate entry", errDuplicatedEntry)
}

func TestApiError_IsDuplicatedEntryErrorFalse(t *testing.T) {
    apiErr := &ApiError{
        Error:      "Another Error",
        StatusCode: http.StatusInternalServerError,
    }
    assert.False(t, apiErr.IsDuplicatedEntryError())
}

func TestApiError_IsDuplicatedEntryErrorTrue(t *testing.T) {
    apiErr := &ApiError{
        Error:      "Duplicate Entry",
        StatusCode: http.StatusInternalServerError,
    }
    assert.True(t, apiErr.IsDuplicatedEntryError())
}
