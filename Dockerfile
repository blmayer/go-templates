FROM golang:1.16 as builder

COPY . /root/

RUN cd /root && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch

COPY --from=builder /root/main /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 27017

CMD ["/main"]