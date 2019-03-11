package shortener

import (
    "fmt"
    "net/http"
    "github.com/emikohmann/url-shortener/src/api/utils"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

const (
    insertMapping      = "insert into url_mapping (hash, value) values (?, ?);"
    selectMapping      = "select value from url_mapping where hash = ?;"
    errMappingNotFound = "mapping not found for hash %s"
)

type URLMessage struct {
    URL string `json:"url"`
}

type URLMapping struct {
    Hash  string
    Value string
}

func (urlMapping *URLMapping) Save() *utils.ApiError {
    if _, err := database.Client.Exec(
        insertMapping,
        urlMapping.Hash,
        urlMapping.Value,
    ); err != nil {
        return &utils.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}

func (urlMapping *URLMapping) Get() *utils.ApiError {
    rows, err := database.Client.Query(
        selectMapping,
        urlMapping.Hash,
    )
    if err != nil {
        return &utils.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    if rows.Next() == false {
        return &utils.ApiError{
            Error:      fmt.Sprintf(errMappingNotFound, urlMapping.Hash),
            StatusCode: http.StatusNotFound,
        }
    }
    if err := rows.Scan(&urlMapping.Value); err != nil {
        return &utils.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}
