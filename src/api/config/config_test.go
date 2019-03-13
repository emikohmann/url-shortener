package config

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "https://jampp.co", SitePrefix)
    assert.EqualValues(t, 5, ShortURLLength)
}
