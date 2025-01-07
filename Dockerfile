FROM golang:1.22.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . . 

WORKDIR /app
RUN go build -o /app/main

CMD ["/app/main"]