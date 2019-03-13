package datadog

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "error sending metric to datadog", errSendingMetricToDD)
    assert.EqualValues(t, "%s.", applicationNamePattern)
    assert.EqualValues(t, "application_name:%s", applicationNameTag)
}

func TestIncrementApplicationMetric(t *testing.T) {
    IncrementApplicationMetric("example_metric", 1)
    time.Sleep(100 * time.Millisecond)
}
