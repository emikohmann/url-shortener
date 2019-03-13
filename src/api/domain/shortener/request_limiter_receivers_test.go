package shortener

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestRequestLimiterConstants(t *testing.T) {
    assert.EqualValues(t, "user_requests", userRequestsTable)
    assert.EqualValues(t, "SELECT request_count FROM %s WHERE user = ? AND minute_id >= ?;", selectBatchUserRequests)
    assert.EqualValues(t, "SELECT request_count FROM %s WHERE user = ? AND minute_id = ? LIMIT 1;", selectUserRequest)
    assert.EqualValues(t, "INSERT INTO %s (user, minute_id, request_count) values (?, ?, ?);", insertUserRequest)
    assert.EqualValues(t, "UPDATE %s SET request_count = ? WHERE user = ? AND minute_id = ?;", updateUserRequest)
    assert.EqualValues(t, "the user %s exceeded the request limit", errUserExceededRequestLimit)
}

func TestURLRequest_CheckExceededRequestLimit(t *testing.T) {
    urlRequest := &URLRequest{
        UserID: "test_user_id",
        URL:    "http://testurl.com",
    }
    if apiErr := urlRequest.CheckExceededRequestLimit(); apiErr != nil {
        t.Error(apiErr)
        return
    }
}

func TestURLRequest_ComputeOneRequest(t *testing.T) {
    urlRequest := &URLRequest{
        UserID: "test_user_id",
        URL:    "http://testurl.com",
    }
    if err := urlRequest.ComputeOneRequest(); err != nil {
        t.Error(err)
        return
    }
}
