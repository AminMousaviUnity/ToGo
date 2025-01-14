# Builder stage
FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

# Install curl and migrate CLI
RUN apk add --no-cache curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz \
    && mv migrate /usr/local/bin/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

COPY .env .
COPY migrations ./migrations

EXPOSE 6666

CMD ["./main"]
