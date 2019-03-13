package shortener

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    domain "github.com/emikohmann/url-shortener/src/api/domain/shortener"
    service "github.com/emikohmann/url-shortener/src/api/services/shortener"
    "github.com/emikohmann/url-shortener/src/api/utils/hashing"
    "github.com/emikohmann/url-shortener/src/api/config"
)

const (
    errInvalidInput    = "invalid input url message"
    headerKeyUserAgent = "User-Agent"
    errShortenURL      = "error in shorten URL"
    errResolveURL      = "error in resolve URL"
    errCountClicks     = "error in count clicks URL"
)

func ShortenURL(c *gin.Context) {
    var input domain.URLRequest

    if err := c.BindJSON(&input); err != nil {
        apiErr := &apierrors.ApiError{
            Error:      errInvalidInput,
            StatusCode: http.StatusBadRequest,
        }
        config.Logger.Println(errShortenURL, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    input.UserID = hashing.MD5(c.GetHeader(headerKeyUserAgent))

    response, apiErr := service.ShortenURL(&input)
    if apiErr != nil {
        config.Logger.Println(errShortenURL, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, response)
}

func ResolveURL(c *gin.Context) {
    var input domain.URLRequest

    if err := c.BindJSON(&input); err != nil {
        apiErr := &apierrors.ApiError{
            Error:      errInvalidInput,
            StatusCode: http.StatusBadRequest,
        }
        config.Logger.Println(errResolveURL, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    input.UserID = hashing.MD5(c.GetHeader(headerKeyUserAgent))

    response, apiErr := service.ResolveURL(&input)
    if apiErr != nil {
        config.Logger.Println(errResolveURL, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, response)
}

func CountClicks(c *gin.Context) {
    var input domain.URLRequest

    if err := c.BindJSON(&input); err != nil {
        apiErr := &apierrors.ApiError{
            Error:      errInvalidInput,
            StatusCode: http.StatusBadRequest,
        }
        config.Logger.Println(errCountClicks, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    input.UserID = hashing.MD5(c.GetHeader(headerKeyUserAgent))

    response, apiErr := service.CountClicks(&input)
    if apiErr != nil {
        config.Logger.Println(errCountClicks, apiErr)
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, response)
}
