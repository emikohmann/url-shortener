package shortener

import (
    "fmt"
    "time"
    "github.com/emikohmann/url-shortener/src/api/utils/vectors"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

const (
    clicksVectorTable       = "clicks_vector"
    selectClicksVector      = "SELECT clicks_count FROM %s WHERE hash = ? AND day_id = ? LIMIT 1;"
    insertClicksVector      = "INSERT INTO %s (hash, day_id, clicks_count) VALUES (?, ?, ?);"
    updateClicksVector      = "UPDATE %s SET clicks_count = ? WHERE hash = ? AND day_id = ?;"
    selectBatchClicksVector = "SELECT day_id, clicks_count FROM %s WHERE hash = ?;"
    dayIDDateFormat         = "20060102"
    dateResponseFormat      = "2006-01-02"
)

func (mapping *Mapping) AggregateVisit() error {
    dayID := vectors.GetDayID(time.Now().UTC())

    rows, err := database.Client.Query(
        fmt.Sprintf(
            selectClicksVector,
            clicksVectorTable,
        ),
        mapping.Hash,
        dayID,
    )
    if err != nil {
        return err
    }
    defer rows.Close()

    exists := rows.Next()
    var clicksCount int64
    if exists {
        if err := rows.Scan(&clicksCount); err != nil {
            return err
        }
    }
    clicksCount++

    switch exists {
    case false:
        _, err := database.Client.Exec(
            fmt.Sprintf(
                insertClicksVector,
                clicksVectorTable,
            ),
            mapping.Hash,
            dayID,
            clicksCount,
        )
        if err != nil {
            return err
        }

    case true:
        _, err := database.Client.Exec(
            fmt.Sprintf(
                updateClicksVector,
                clicksVectorTable,
            ),
            clicksCount,
            mapping.Hash,
            dayID,
        )
        if err != nil {
            return err
        }
    }

    return nil
}

func (mapping *Mapping) CountClicks() (*ClicksCounterResponse, error) {
    result := make(ClicksCounterResponse, 0)

    rows, err := database.Client.Query(
        fmt.Sprintf(
            selectBatchClicksVector,
            clicksVectorTable,
        ),
        mapping.Hash,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var currentDayID string
        var currentClicksCount int64

        rows.Scan(&currentDayID, &currentClicksCount)

        date, err := time.Parse(dayIDDateFormat, currentDayID)
        if err != nil {
            return nil, err
        }

        result = append(result, ClicksCounter{
            Date:   date.Format(dateResponseFormat),
            Visits: currentClicksCount,
        })
    }

    return &result, nil
}
