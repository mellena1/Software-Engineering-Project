FROM golang:1.11.5 as builder

WORKDIR /api
RUN go get -u github.com/swaggo/swag/cmd/swag
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init -s docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o backend

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /api/backend /backend

ENTRYPOINT ["/backend"]
