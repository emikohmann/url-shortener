package shortener

import (
    "fmt"
    "sync"
    "net/http"
    "database/sql"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

const (
    insertURLMapping   = "INSERT INTO url_mapping (url, hash) VALUES (?, ?);"
    insertHashMapping  = "INSERT INTO hash_mapping (hash, url) VALUES (?, ?);"
    selectURLFromHash  = "SELECT url FROM hash_mapping WHERE hash = ?;"
    selectHashFromURL  = "SELECT hash FROM url_mapping WHERE url = ?;"
    errMappingNotFound = "mapping not found for %s"
)

// save mapping in db and check consistency
func (mapping *Mapping) Save() *apierrors.ApiError {
    transaction, err := database.Client.Begin()
    if err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }

    out := make(chan *apierrors.ApiError, 2)
    defer close(out)

    // save url and hash mapping both in parallel
    var group sync.WaitGroup
    group.Add(2)
    go mapping.AsyncSaveURLMapping(transaction, out, &group)
    go mapping.AsyncSaveHashMapping(transaction, out, &group)
    group.Wait()

    for i := 0; i < 2; i++ {
        executionErr := <-out
        if executionErr != nil {
            // if any operation fails, rollback
            if err := transaction.Rollback(); err != nil {
                return &apierrors.ApiError{
                    Error:      err.Error(),
                    StatusCode: http.StatusInternalServerError,
                }
            }
            return executionErr
        }
    }

    if err := transaction.Commit(); err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}

func (mapping *Mapping) AsyncSaveURLMapping(transaction *sql.Tx, out chan *apierrors.ApiError, group *sync.WaitGroup) {
    defer group.Done()
    if apiErr := mapping.SaveURLMapping(transaction); apiErr != nil {
        out <- apiErr
        return
    }
    out <- nil
}

func (mapping *Mapping) AsyncSaveHashMapping(transaction *sql.Tx, out chan *apierrors.ApiError, group *sync.WaitGroup) {
    defer group.Done()
    if apiErr := mapping.SaveHashMapping(transaction); apiErr != nil {
        out <- apiErr
        return
    }
    out <- nil
}

func (mapping *Mapping) SaveURLMapping(transaction *sql.Tx) *apierrors.ApiError {
    if _, err := transaction.Exec(
        insertURLMapping,
        mapping.URL,
        mapping.Hash,
    ); err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}

func (mapping *Mapping) SaveHashMapping(transaction *sql.Tx) *apierrors.ApiError {
    if _, err := transaction.Exec(
        insertHashMapping,
        mapping.Hash,
        mapping.URL,
    ); err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}

func (mapping *Mapping) GetHashFromURL() *apierrors.ApiError {
    rows, err := database.Client.Query(
        selectHashFromURL,
        mapping.URL,
    )
    if err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    if rows.Next() == false {
        return &apierrors.ApiError{
            Error:      fmt.Sprintf(errMappingNotFound, mapping.URL),
            StatusCode: http.StatusNotFound,
        }
    }
    if err := rows.Scan(&mapping.Hash); err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}

func (mapping *Mapping) GetURLFromHash() *apierrors.ApiError {
    rows, err := database.Client.Query(
        selectURLFromHash,
        mapping.Hash,
    )
    if err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    if rows.Next() == false {
        return &apierrors.ApiError{
            Error:      fmt.Sprintf(errMappingNotFound, mapping.Hash),
            StatusCode: http.StatusNotFound,
        }
    }
    if err := rows.Scan(&mapping.URL); err != nil {
        return &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
    }
    return nil
}
