# API Porticos - dockerfile
FROM golang:1.25.5-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/api ./cmd/api

FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/api /app/api

EXPOSE 3200

CMD ["/app/api"]
