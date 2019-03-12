package hashing

import (
    "fmt"
    "time"
    "regexp"
    "net/http"
    "math/rand"
    "github.com/emikohmann/url-shortener/src/api/config"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
)

const (
    errHashNotFound = "hash not found in url"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var (
    letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func RandomString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func ExtractHash(url string) (string, *apierrors.ApiError) {
    reg, err := regexp.Compile(fmt.Sprintf("%s/%s", config.SitePrefix, `([a-zA-Z0-9]+)$`))
    if err != nil {
        return "", &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusBadRequest,
        }
    }
    if result := reg.FindStringSubmatch(url); len(result) > 1 {
        return result[1], nil
    }
    return "", &apierrors.ApiError{
        Error:      errHashNotFound,
        StatusCode: http.StatusBadRequest,
    }
}
