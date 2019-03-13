package config

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "url-shortener", ApplicationName)
    assert.EqualValues(t, "127.0.0.1:8125", DatadogMetricAddress)
    assert.EqualValues(t, "https://jampp.co", SitePrefix)
    assert.EqualValues(t, 5, ShortURLLength)
    assert.EqualValues(t, 10, RateLimiterMaxRequestCount)
    assert.EqualValues(t, 1*time.Hour, RateLimiteraxRequestTime)
}
