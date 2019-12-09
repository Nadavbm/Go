FROM golang:1.11.4-stretch AS builder

COPY src /go/src/envinter

WORKDIR /go/src/envinter

ENV CGO_ENABLED=0
ENV GO111MODULE=off
ENV GOOS=linux

RUN go build -o /go/bin/envInterpreter

FROM alpine:latest

ENV URL1="http://localhost.local"
ENV PORT=8089

COPY --from=builder /go/bin/envInterpreter  /app/envInterpreter 

CMD ["/app/envInterpreter"]

