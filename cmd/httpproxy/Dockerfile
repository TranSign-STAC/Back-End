FROM golang:1.15.7

ENV GO111MODULE=on

ADD . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o httpproxy cmd/httpproxy/main.go

EXPOSE 8080

ENTRYPOINT [ "./httpproxy" ]