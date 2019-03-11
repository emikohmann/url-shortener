package app

import (
    "github.com/gin-gonic/gin"
    "github.com/emikohmann/url-shortener/src/api/controllers/ping"
    "github.com/emikohmann/url-shortener/src/api/controllers/shortener"
)

const (
    port = ":8080"
)

var (
    router = gin.Default()
)

func mapRoutes() {
    // health check
    router.GET("/ping", ping.Ping)

    // shorten url
    router.POST("/shorten", shortener.ShortenURL)

    // resolve url
    router.POST("/resolve", shortener.ResolveURL)
}

func run() {
    router.Run(port)
}
