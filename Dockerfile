FROM golang:1.15.7

ENV GO111MODULE=on

WORKDIR /app

COPY cmd/ cmd/
COPY configs/ configs/
COPY gen/ gen/
COPY go.mod .
COPY go.sum .
COPY build.sh .
COPY run.sh .

RUN chmod +x build.sh
RUN chmod +x run.sh
RUN go mod download
RUN sh build.sh

EXPOSE 8000 8080

ENTRYPOINT [ "sh", "run.sh" ]