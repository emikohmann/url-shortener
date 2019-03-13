FROM golang:alpine
WORKDIR /go/src/github.com/emikohmann/url-shortener
COPY . .
RUN go build -o /go/bin/url-shortener src/api/main.go

FROM scratch
COPY --from=build /go/bin/url-shortener /go/bin/url-shortener
ENTRYPOINT ["/go/bin/url-shortener"]