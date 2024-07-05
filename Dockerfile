# Stage 1: Build the Go binary
FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env

WORKDIR /app/cmd/server
RUN go build -o /app/main .

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

ENTRYPOINT ["/app/main"]
