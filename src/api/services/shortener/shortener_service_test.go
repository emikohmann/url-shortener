package shortener

import (
    "os"
    "bytes"
    "testing"
    "net/http"
    "github.com/stretchr/testify/assert"
    "github.com/emikohmann/url-shortener/src/api/domain/shortener"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

func TestMain(m *testing.M) {
    database.Connect()
    os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
    assert.EqualValues(t, 3000, maxURLSize)
    assert.EqualValues(t, "url is too long", errURLTooLong)
    assert.EqualValues(t, "%s/%s", urlPattern)
    assert.EqualValues(t, "error vectorizing visit", errVectorizingVisit)
}

func TestShortenURLTooLong(t *testing.T) {
    var buff bytes.Buffer
    for i := 0; i < 1000; i++ {
        buff.WriteString("too_much")
    }
    _, apiErr := ShortenURL(&shortener.URLRequest{
        URL: buff.String(),
    })
    assert.NotNil(t, apiErr)
    assert.EqualValues(t, http.StatusBadRequest, apiErr.StatusCode)
    assert.EqualValues(t, "url is too long", apiErr.Error)
}

func TestShortenURL(t *testing.T) {
    response, apiErr := ShortenURL(&shortener.URLRequest{
        URL: "http://www.example.com",
    })
    if apiErr != nil {
        t.Error(apiErr)
        return
    }
    assert.NotEmpty(t, response)
}

func TestResolveURLInvalidHash(t *testing.T) {
    _, apiErr := ResolveURL(&shortener.URLRequest{
        URL: "https://jampp.co/invalid/hash",
    })
    assert.NotNil(t, apiErr)
    assert.EqualValues(t, http.StatusBadRequest, apiErr.StatusCode)
    assert.EqualValues(t, "hash not found in url", apiErr.Error)
}

func TestResolveURL(t *testing.T) {
    _, err := database.Client.Exec(
        "delete from hash_mapping where hash = 'testhash'",
    )
    if err != nil {
        t.Error(err)
        return
    }
    _, err = database.Client.Exec(
        "delete from url_mapping where url = 'https://other.com/test_url'",
    )
    if err != nil {
        t.Error(err)
        return
    }
    _, err = database.Client.Exec(
        "insert into hash_mapping (hash, url) values ('testhash', 'https://other.com/test_url')",
    )
    if err != nil {
        t.Error(err)
        return
    }
    _, err = database.Client.Exec(
        "insert into url_mapping (url, hash) values ('https://other.com/test_url', 'testhash')",
    )
    if err != nil {
        t.Error(err)
        return
    }
    response, apiErr := ResolveURL(&shortener.URLRequest{
        URL: "https://jampp.co/testhash",
    })
    if apiErr != nil {
        t.Error(apiErr)
        return
    }
    assert.EqualValues(t, "https://other.com/test_url", response.ResolvedURL)
}

func TestCountClicksInvalidHash(t *testing.T) {
    _, apiErr := CountClicks(&shortener.URLRequest{
        URL: "https://jampp.co/invalid/hash",
    })
    assert.NotNil(t, apiErr)
    assert.EqualValues(t, http.StatusBadRequest, apiErr.StatusCode)
    assert.EqualValues(t, "hash not found in url", apiErr.Error)
}

func TestCountClicks(t *testing.T) {
    response, apiErr := CountClicks(&shortener.URLRequest{
        URL: "https://jampp.co/newhash",
    })
    if apiErr != nil {
        t.Error(apiErr)
        return
    }
    assert.NotNil(t, response)
    assert.Empty(t, *response)
}
