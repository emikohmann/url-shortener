package shortener

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    domain "github.com/emikohmann/url-shortener/src/api/domain/shortener"
    service "github.com/emikohmann/url-shortener/src/api/services/shortener"
)

const (
    errInvalidInput = "invalid input url message"
)

func ShortenURL(c *gin.Context) {
    var input domain.URLRequest

    if err := c.BindJSON(&input); err != nil {
        apiErr := &apierrors.ApiError{
            Error:      errInvalidInput,
            StatusCode: http.StatusBadRequest,
        }
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    response, apiErr := service.ShortenURL(&input)
    if apiErr != nil {
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
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    response, apiErr := service.ResolveURL(&input)
    if apiErr != nil {
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
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    response, apiErr := service.CountClicks(&input)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, response)
}
