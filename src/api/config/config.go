package config

import (
    "log"
    "time"
    "os"
)

const (
    ApplicationName            = "url-shortener"
    DatadogMetricAddress       = "127.0.0.1:8125"
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
