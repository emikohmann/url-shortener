#### URL Shortener

To run this project in your computer, load database connection config in `load_env.sh` and execute:

```bash
cd src/api
source load_env.sh
```

Then, build and run the project:

```bash
go build
./api
```

Or just:

```bash
go run main.go
```

---

#### Available endpoints

* **POST** `/shorten` receives an URL to be shortened.

```json
{
    "url": "https://www.facebook.com"
}
```

This endpoint is limited to (configurable) 10 request in last hour per user. if user exceeds the limit, he will receive a response code 429.

Example response:

```json
{
    "short_url": "https://jampp.co/sH4m3"
}
```

* **POST** `/resolve` receives an short URL to be resolved.

```json
{
    "url": "https://jampp.co/sH4m3"
}
```

Example response:

```json
{
    "resolved_url": "https://www.facebook.com"
}
```

* **POST** `/clicks` compute clicks counter per hour for a particular shor URL.

```json
{
    "url": "https://jampp.co/sH4m3"
}
```

Example response:

```json
[
    {
        "date": "2019-03-12",
        "clicks": 12
    },
    {
        "date": "2019-03-13",
        "clicks": 8
    }
]
```