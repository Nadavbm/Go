FROM golang:1.12.10 AS builder

RUN mkdir /src

COPY src/main.go /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/simpleApiServer

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /go/bin/simpleApiServer /app/simpleApiServer

EXPOSE 8080
STOPSIGNAL SIGTERM

CMD ["/app/simpleApiServer"]

