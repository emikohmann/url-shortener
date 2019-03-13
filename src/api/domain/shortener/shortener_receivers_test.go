package shortener

import (
    "sync"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/emikohmann/url-shortener/src/api/utils/apierrors"
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

func TestShortenerConstants(t *testing.T) {
    assert.EqualValues(t, "url_mapping", tableURLMapping)
    assert.EqualValues(t, "hash_mapping", tableHashMapping)
    assert.EqualValues(t, "INSERT INTO %s (url, hash) VALUES (?, ?);", insertURLMapping)
    assert.EqualValues(t, "INSERT INTO %s (hash, url) VALUES (?, ?);", insertHashMapping)
    assert.EqualValues(t, "SELECT url FROM %s WHERE hash = ?;", selectURLFromHash)
    assert.EqualValues(t, "SELECT hash FROM %s WHERE url = ?;", selectHashFromURL)
    assert.EqualValues(t, "mapping not found for %s", errMappingNotFound)
}

func TestMapping_Save(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    if apiErr := mapping.Save(); apiErr != nil {
        if !apiErr.IsDuplicatedEntryError() {
            t.Error(apiErr)
            return
        }
    }
}

func TestMapping_AsyncSaveURLMapping(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    tx, err := database.Client.Begin()
    if err != nil {
        t.Error(err)
        return
    }
    out := make(chan *apierrors.ApiError, 1)
    var wg sync.WaitGroup
    wg.Add(1)
    go mapping.AsyncSaveURLMapping(tx, out, &wg)
    wg.Wait()
    executionErr := <-out
    if executionErr != nil {
        if err := tx.Rollback(); err != nil {
            t.Error(err)
            return
        }
        if !executionErr.IsDuplicatedEntryError() {
            t.Error(executionErr)
            return
        }
    }
}

func TestMapping_AsyncSaveHashMapping(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    tx, err := database.Client.Begin()
    if err != nil {
        t.Error(err)
        return
    }
    out := make(chan *apierrors.ApiError, 1)
    var wg sync.WaitGroup
    wg.Add(1)
    go mapping.AsyncSaveURLMapping(tx, out, &wg)
    wg.Wait()
    executionErr := <-out
    if executionErr != nil {
        if err := tx.Rollback(); err != nil {
            t.Error(err)
            return
        }
        if !executionErr.IsDuplicatedEntryError() {
            t.Error(executionErr)
            return
        }
    }
}

func TestMapping_SaveHashMapping(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    tx, err := database.Client.Begin()
    if err != nil {
        t.Error(err)
        return
    }
    if apiErr := mapping.SaveHashMapping(tx); apiErr != nil {
        if !apiErr.IsDuplicatedEntryError() {
            t.Error(apiErr)
            return
        }
    }
}

func TestMapping_SaveURLMapping(t *testing.T) {
    mapping := &Mapping{
        URL:  "http://www.facebook.com",
        Hash: "https://jampp.co/abc12",
    }
    tx, err := database.Client.Begin()
    if err != nil {
        t.Error(err)
        return
    }
    if apiErr := mapping.SaveURLMapping(tx); apiErr != nil {
        if !apiErr.IsDuplicatedEntryError() {
            t.Error(apiErr)
            return
        }
    }
}

func TestMapping_GetHashFromURL(t *testing.T) {
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
    mapping := &Mapping{
        URL: "https://other.com/test_url",
    }
    if apiErr := mapping.GetHashFromURL(); apiErr != nil {
        t.Error(apiErr)
        return
    }
    assert.EqualValues(t, "testhash", mapping.Hash)
}

func TestMapping_GetURLFromHash(t *testing.T) {
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
    mapping := &Mapping{
        Hash: "testhash",
    }
    if apiErr := mapping.GetURLFromHash(); apiErr != nil {
        t.Error(apiErr)
        return
    }
    assert.EqualValues(t, "https://other.com/test_url", mapping.URL)
}
