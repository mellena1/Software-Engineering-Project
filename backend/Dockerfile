FROM golang:1.11.5-alpine

WORKDIR /api
RUN apk add git gcc musl-dev
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api

ENTRYPOINT ["./api"]
