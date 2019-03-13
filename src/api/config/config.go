package config

import "time"

const (
    SitePrefix                 = "https://jampp.co"
    ShortURLLength             = 5
    RateLimiterMaxRequestCount = 10
    RateLimiteraxRequestTime   = 1 * time.Hour
)
