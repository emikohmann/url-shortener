package vectors

import "time"

const (
    dayYearModifier  int64 = 10000
    dayMonthModifier int64 = 100
    dayDayModifier   int64 = 1

    minuteYearModifier   int64 = 100000000
    minuteMonthModifier  int64 = 1000000
    minuteDayModifier    int64 = 10000
    minuteHourModifier   int64 = 100
    minuteMinuteModifier int64 = 1
)

func GetDayID(date time.Time) int64 {
    year, month, day := date.Date()

    return func(values ...int64) int64 {
        var total int64
        for _, v := range values {
            total += v
        }
        return total
    }(
        int64(year)*dayYearModifier,
        int64(month)*dayMonthModifier,
        int64(day)*dayDayModifier,
    )
}

func GetMinuteID(date time.Time) int64 {
    year, month, day := date.Date()
    hour := date.Hour()
    minute := date.Minute()

    return func(values ...int64) int64 {
        var total int64
        for _, v := range values {
            total += v
        }
        return total
    }(
        int64(year)*minuteYearModifier,
        int64(month)*minuteMonthModifier,
        int64(day)*minuteDayModifier,
        int64(hour)*minuteHourModifier,
        int64(minute)*minuteMinuteModifier,
    )
}
