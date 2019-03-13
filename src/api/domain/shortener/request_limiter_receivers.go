package shortener

import (
    "fmt"
    "time"
    "net/http"
    "github.com/emikohmann/url-shortener/src/api/config"
    "github.com/emikohmann/url-shortener/src/api/utils/vectors"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

const (
    userRequestsTable           = "user_requests"
    selectBatchUserRequests     = "SELECT request_count FROM %s WHERE user = ? AND minute_id >= ?;"
    selectUserRequest           = "SELECT request_count FROM %s WHERE user = ? AND minute_id = ? LIMIT 1;"
    insertUserRequest           = "INSERT INTO %s (user, minute_id, request_count) values (?, ?, ?);"
    updateUserRequest           = "UPDATE %s SET request_count = ? WHERE user = ? AND minute_id = ?;"
    errUserExceededRequestLimit = "the user %s exceeded the request limit"
)

func (urlRequest *URLRequest) CheckExceededRequestLimit() *apierrors.ApiError {
    rows, err := database.Client.Query(
        fmt.Sprintf(
            selectBatchUserRequests,
            userRequestsTable,
        ),
        urlRequest.UserID,
        vectors.GetMinuteID(
            time.Now().UTC().Add(
                -config.RateLimiterMaxRequestCount*config.RateLimiteraxRequestTime,
            ),
        ),
    )
    if err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    defer rows.Close()

    var totalRequest int64
    for rows.Next() {
        var currentClicksCount int64
        if err := rows.Scan(&currentClicksCount); err != nil {
            return &apierrors.ApiError{
                Error:      err.Error(),
                StatusCode: http.StatusInternalServerError,
            }
        }
        totalRequest += currentClicksCount
    }
    if totalRequest >= config.RateLimiterMaxRequestCount {
        return &apierrors.ApiError{
            Error: fmt.Sprintf(
                errUserExceededRequestLimit,
                urlRequest.UserID,
            ),
            StatusCode: http.StatusTooManyRequests,
        }
    }
    return nil
}

func (urlRequest *URLRequest) ComputeOneRequest() error {
    minuteID := vectors.GetMinuteID(time.Now().UTC())

    rows, err := database.Client.Query(
        fmt.Sprintf(
            selectUserRequest,
            userRequestsTable,
        ),
        urlRequest.UserID,
        minuteID,
    )
    if err != nil {
        return err
    }
    defer rows.Close()

    exists := rows.Next()
    var requestCount int64
    if exists {
        if err := rows.Scan(&requestCount); err != nil {
            return err
        }
    }
    requestCount++

    switch exists {
    case false:
        _, err := database.Client.Exec(
            fmt.Sprintf(
                insertUserRequest,
                userRequestsTable,
            ),
            urlRequest.UserID,
            minuteID,
            requestCount,
        )
        if err != nil {
            return err
        }

    case true:
        _, err := database.Client.Exec(
            fmt.Sprintf(
                updateUserRequest,
                userRequestsTable,
            ),
            requestCount,
            urlRequest.UserID,
            minuteID,
        )
        if err != nil {
            return err
        }
    }

    return nil
}
