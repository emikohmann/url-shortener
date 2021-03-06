package vectors

import (
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, 10000, dayYearModifier)
    assert.EqualValues(t, 100, dayMonthModifier)
    assert.EqualValues(t, 1, dayDayModifier)

    assert.EqualValues(t, 100000000, minuteYearModifier)
    assert.EqualValues(t, 1000000, minuteMonthModifier)
    assert.EqualValues(t, 10000, minuteDayModifier)
    assert.EqualValues(t, 100, minuteHourModifier)
    assert.EqualValues(t, 1, minuteMinuteModifier)
}

func TestGetDayID(t *testing.T) {
    date, err := time.Parse("2006-01-02T15:04:05.000Z", "2014-11-12T11:45:26.371Z")
    if err != nil {
        t.Error(err)
        return
    }
    assert.EqualValues(t, 20141112, GetDayID(date))
}

func TestGetMinuteID(t *testing.T) {
    date, err := time.Parse("2006-01-02T15:04:05.000Z", "2014-11-12T11:45:26.371Z")
    if err != nil {
        t.Error(err)
        return
    }
    assert.EqualValues(t, 201411121145, GetMinuteID(date))
}
