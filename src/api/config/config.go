package config

import (
    "log"
    "time"
    "os"
)

const (
    SitePrefix                 = "https://jampp.co"
    ShortURLLength             = 5
    RateLimiterMaxRequestCount = 10
    RateLimiteraxRequestTime   = 1 * time.Hour
)

var (
    Logger log.Logger
)

func init() {
    Logger.SetOutput(os.Stdout)
}
