FROM golang:1.11.5-alpine

WORKDIR /api
RUN apk add git gcc musl-dev
RUN go get -u github.com/swaggo/swag/cmd/swag
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init -s docs
RUN go build -o api

ENTRYPOINT ["./api"]
