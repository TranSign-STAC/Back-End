FROM golang:1.15.7

ENV GO111MODULE=on

WORKDIR /app

COPY cmd/ cmd/
COPY configs/ configs/
COPY gen/ gen/
COPY go.mod .
COPY go.sum .
COPY run.sh .

RUN go mod download
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o server cmd/server/main.go
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o httpproxy cmd/httpproxy/main.go

RUN chmod +x run.sh

EXPOSE 8000 8080

ENTRYPOINT [ "sh", "run.sh" ]