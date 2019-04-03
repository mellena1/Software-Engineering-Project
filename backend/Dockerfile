FROM golang:1.11.5 as builder

WORKDIR /api
RUN curl -L -o swag.tar.gz https://github.com/swaggo/swag/releases/download/v1.4.1/swag_1.4.1_Linux_x86_64.tar.gz && \
    tar -xvf swag.tar.gz && \
    mv swag /bin/swag
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init -s docs
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o backend

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /api/backend /backend

ENTRYPOINT ["/backend"]
