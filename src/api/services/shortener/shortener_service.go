package shortener

import (
    "fmt"
    "net/http"
    "github.com/emikohmann/url-shortener/src/api/utils"
    "github.com/emikohmann/url-shortener/src/api/domain/shortener"
)

const (
    maxURLSize    = 3000
    errURLTooLong = "url is to long"
    urlPattern    = "http://jampp.go/%s"
)

func ShortenURL(input *shortener.URLMessage) (*shortener.URLMessage, *utils.ApiError) {
    if len(input.URL) > maxURLSize {
        return nil, &utils.ApiError{
            Error:      errURLTooLong,
            StatusCode: http.StatusBadRequest,
        }
    }

    mapping := shortener.URLMapping{
        Hash:  utils.MD5(input.URL),
        Value: input.URL,
    }

    if apiErr := mapping.Save(); apiErr != nil {
        return nil, apiErr
    }

    return &shortener.URLMessage{
        URL: fmt.Sprintf(urlPattern, mapping.Hash),
    }, nil
}

func ResolveURL(input *shortener.URLMessage) (*shortener.URLMessage, *utils.ApiError) {
    hash, apiErr := utils.ExtractHash(input.URL)
    if apiErr != nil {
        return nil, apiErr
    }

    mapping := shortener.URLMapping{
        Hash: hash,
    }

    if apiErr := mapping.Get(); apiErr != nil {
        return nil, apiErr
    }

    return &shortener.URLMessage{
        URL: mapping.Value,
    }, nil
}
