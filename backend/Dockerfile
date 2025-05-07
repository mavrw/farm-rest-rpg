# backend/Dockerfile
FROM golang:1.24 AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o server ./cmd/api

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache libc6-compat
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
