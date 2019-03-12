package shortener

type URLRequest struct {
    URL string `json:"url"`
}

type ShortenURLResponse struct {
    ShortURL string `json:"short_url"`
}

type ResolveURLResponse struct {
    ResolvedURL string `json:"resolved_url"`
}

type Mapping struct {
    URL  string
    Hash string
}

type ClicksCounterResponse []ClicksCounter

type ClicksCounter struct {
    Date   string `json:"date"`
    Visits int64  `json:"visits"`
}
