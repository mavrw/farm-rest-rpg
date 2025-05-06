# backend/Dockerfile
FROM golang:1.24 AS builder

WORKDIR /app

COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o server ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
