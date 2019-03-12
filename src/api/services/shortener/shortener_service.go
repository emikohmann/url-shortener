package shortener

import (
    "fmt"
    "net/http"
    "github.com/emikohmann/url-shortener/src/api/config"
    "github.com/emikohmann/url-shortener/src/api/utils/hashing"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    "github.com/emikohmann/url-shortener/src/api/domain/shortener"
)

const (
    maxURLSize          = 3000
    errURLTooLong       = "url is too long"
    urlPattern          = "%s/%s"
    errVectorizingVisit = "error vectorizing visit"
)

func ShortenURL(input *shortener.URLRequest) (*shortener.ShortenURLResponse, *apierrors.ApiError) {
    if len(input.URL) > maxURLSize {
        return nil, &apierrors.ApiError{
            Error:      errURLTooLong,
            StatusCode: http.StatusBadRequest,
        }
    }

    mapping := shortener.Mapping{
        URL: input.URL,
    }

    // try to get existing record
    if apiErr := mapping.GetHashFromURL(); apiErr != nil {
        if apiErr.StatusCode != http.StatusNotFound {
            return nil, apiErr
        }

        // is new url, shorten
        for {
            mapping.Hash = hashing.RandomString(config.ShortURLLength)
            if apiErr := mapping.Save(); apiErr != nil {
                if !apiErr.IsDuplicatedEntryError() {
                    return nil, apiErr
                }
                // continue if collides with existing hash
                continue
            }
            // break if hash ok
            break
        }
    }

    return &shortener.ShortenURLResponse{
        ShortURL: fmt.Sprintf(urlPattern, config.SitePrefix, mapping.Hash),
    }, nil
}

func ResolveURL(input *shortener.URLRequest) (*shortener.ResolveURLResponse, *apierrors.ApiError) {
    hash, apiErr := hashing.ExtractHash(input.URL)
    if apiErr != nil {
        return nil, apiErr
    }

    mapping := shortener.Mapping{
        Hash: hash,
    }

    if apiErr := mapping.GetURLFromHash(); apiErr != nil {
        return nil, apiErr
    }

    go func() {
        if err := mapping.AggregateVisit(); err != nil {
            fmt.Println(errVectorizingVisit, err)
            return
        }
    }()

    return &shortener.ResolveURLResponse{
        ResolvedURL: mapping.URL,
    }, nil
}

func CountClicks(input *shortener.URLRequest) (*shortener.ClicksCounterResponse, *apierrors.ApiError) {
    hash, apiErr := hashing.ExtractHash(input.URL)
    if apiErr != nil {
        return nil, apiErr
    }

    mapping := shortener.Mapping{
        Hash: hash,
    }

    response, err := mapping.CountClicks()
    if err != nil {
        apiErr := &apierrors.ApiError{
            Error:      err.Error(),
            StatusCode: http.StatusInternalServerError,
        }
        return nil, apiErr
    }
    return response, nil
}
