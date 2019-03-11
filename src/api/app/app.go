package app

import (
    "github.com/emikohmann/url-shortener/src/api/clients/database"
)

func StartApp() {
    // init services
    database.Connect()

    // map urls to controllers
    mapRoutes()

    // run app
    run()
}
