package app

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, ":8080", port)
}
