package vectors

import "time"

const (
    dayYearModifier  int64 = 10000
    dayMonthModifier int64 = 100
    dayDayModifier   int64 = 1
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
