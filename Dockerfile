FROM golang:alpine

ENV DB_HOST=localhost:3306
ENV DB_USERNAME=root
ENV DB_PASSWORD=your_mysql_pwd
ENV DB_SCHEMA=shortener

WORKDIR /go/src/github.com/emikohmann/url-shortener
COPY . .
RUN go build -o /go/bin/url-shortener src/api/main.go

FROM scratch
COPY --from=build /go/bin/url-shortener /go/bin/url-shortener
ENTRYPOINT ["/go/bin/url-shortener"]