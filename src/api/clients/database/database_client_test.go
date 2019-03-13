package database

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "Database connected", infDatabaseConnected)
    assert.EqualValues(t, "Error connecting database", errConnectingDatabase)
}

func TestConnect(t *testing.T) {
    Connect()
}
