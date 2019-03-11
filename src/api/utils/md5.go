package utils

import (
    "fmt"
    "regexp"
    "net/http"
    "crypto/md5"
)

const (
    errHashNotFound = "hash not found in url"
)

func MD5(input string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func ExtractHash(url string) (string, *ApiError) {
    reg, err := regexp.Compile(`http://jampp.go/([a-zA-Z0-9]+)$`)
    if err != nil {
        return "", &ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    if result := reg.FindStringSubmatch(url); len(result) > 1 {
        return result[1], nil
    }
    return "", &ApiError{
        Error:      errHashNotFound,
        StatusCode: http.StatusNotFound,
    }
}
