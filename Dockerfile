FROM golang:1.13.1-strech AS builder

ENV CGO_ENABLED=0

RUN go build -o /go/bin/appName path/to/project

FROM alpine:latest

COPY --from=builder /go/bin/appName /app/appName

CMD ["/app/appName"]

