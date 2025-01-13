FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

EXPOSE 6666

CMD ["./main"]
