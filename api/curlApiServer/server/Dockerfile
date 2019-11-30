# Server build
FROM golang:1.11 as builder

RUN go get -v -u github.com/gorilla/mux

COPY . /go/src/server 

WORKDIR /go/src/server/api 

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /go/bin/apiServer

# Running server on alpine
FROM alpine:latest

COPY --from=builder /go/bin/apiServer .

RUN apk add --no-cache libc6-compat && \
    chmod 755 /apiServer

EXPOSE 8081

CMD ["./apiServer"]
