package hashing

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
    assert.EqualValues(t, "hash not found in url", errHashNotFound)
}

func TestVars(t *testing.T) {
    assert.EqualValues(t, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"), letterRunes)
}

func TestRandomString(t *testing.T) {
    assert.EqualValues(t, 1, len(RandomString(1)))
    assert.EqualValues(t, 2, len(RandomString(2)))
    assert.EqualValues(t, 3, len(RandomString(3)))
    assert.EqualValues(t, 4, len(RandomString(4)))
    assert.EqualValues(t, 5, len(RandomString(5)))
    assert.EqualValues(t, 6, len(RandomString(6)))
    assert.EqualValues(t, 7, len(RandomString(7)))
    assert.EqualValues(t, 8, len(RandomString(8)))
    assert.EqualValues(t, 9, len(RandomString(9)))
}

func TestExtractHash(t *testing.T) {
    for i := 0; i < 100; i++ {
        randomHash := RandomString(5)
        hash, apiErr := ExtractHash(fmt.Sprintf("https://jampp.co/%s", randomHash))
        if apiErr != nil {
            t.Error(apiErr)
            return
        }
        assert.EqualValues(t, randomHash, hash)
    }
}
