FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM alpine:3.23

RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/login /app/login
ENTRYPOINT ["/app/login"]