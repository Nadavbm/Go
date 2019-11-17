FROM golang:1.12.10 AS builder

RUN mkdir /src

COPY src/main.go /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/simpleHttpServer

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /go/bin/simpleHttpServer /app/simpleHttpServer

EXPOSE 8080
STOPSIGNAL SIGTERM

CMD ["/app/simpleHttpServer"]

