# Client build
FROM golang:1.11 as builder

#RUN go build -tags netgo -a -v

COPY . /go/src/client 

WORKDIR /go/src/client/bin 

ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux

RUN go build -o /go/bin/apiClient

# Running server on alpine
FROM alpine:latest

ADD scripts/ /scripts

COPY --from=builder /go/bin/apiClient .

RUN apk add --no-cache libc6-compat && \
    chmod 755 /apiClient && \
    chmod 755 /scripts/cron.sh && \
    echo "* *   *   *   *   cd / && /apiClient" >> /etc/crontabs/root && \
    echo "* *   *   *   *   /scripts/cron.sh" >> /etc/crontabs/root

ENV POST_URL="http://192.168.178.91:8081/text"

CMD crond -l 2 -f