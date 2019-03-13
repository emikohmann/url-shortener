package shortener

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

func TestMain(m *testing.M) {
    database.Connect()
    os.Exit(m.Run())
}

func TestClicksVectorConstants(t *testing.T) {
    assert.EqualValues(t, "clicks_vector", clicksVectorTable)
    assert.EqualValues(t, "SELECT clicks_count FROM %s WHERE hash = ? AND day_id = ? LIMIT 1;", selectClicksVector)
    assert.EqualValues(t, "INSERT INTO %s (hash, day_id, clicks_count) VALUES (?, ?, ?);", insertClicksVector)
    assert.EqualValues(t, "UPDATE %s SET clicks_count = ? WHERE hash = ? AND day_id = ?;", updateClicksVector)
    assert.EqualValues(t, "SELECT day_id, clicks_count FROM %s WHERE hash = ?;", selectBatchClicksVector)
    assert.EqualValues(t, "20060102", dayIDDateFormat)
    assert.EqualValues(t, "2006-01-02", dateResponseFormat)
}

func TestMapping_AggregateVisit(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    if err := mapping.AggregateVisit(); err != nil {
        t.Error(err)
        return
    }
}

func TestMapping_CountClicks(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    if err := mapping.AggregateVisit(); err != nil {
        t.Error(err)
        return
    }
    response, err := mapping.CountClicks()
    if err != nil {
        t.Error(err)
        return
    }
    assert.NotNil(t, response)
    assert.NotEmpty(t, *response)
}
