FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go build -o /app/main .

FROM alpine:latest  

COPY --from=builder /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]
