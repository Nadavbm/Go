FROM alpine:3.8

WORKDIR /usr/local

COPY webservice .

EXPOSE 8080

STOPSIGNAL SIGTERM

CMD ["./webservice"]
