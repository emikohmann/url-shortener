package shortener

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "invalid input url message", errInvalidInput)
    assert.EqualValues(t, "User-Agent", headerKeyUserAgent)
    assert.EqualValues(t, "error in shorten URL", errShortenURL)
    assert.EqualValues(t, "error in resolve URL", errResolveURL)
    assert.EqualValues(t, "error in count clicks URL", errCountClicks)
    assert.EqualValues(t, "shorten_url_requests", metricShortenURLRequests)
    assert.EqualValues(t, "resolve_url_requests", metricResolveURLRequests)
    assert.EqualValues(t, "count_clicks_requests", metricCountClicksRequests)
    assert.EqualValues(t, "final_status:%d", metricTagFinalStatus)
}
