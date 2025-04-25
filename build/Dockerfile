# build
FROM golang:1.24.1-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git make

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN make install-swag swag

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/main.go

# app
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .
COPY --from=builder /app/docs .
COPY --from=builder /app/config/*.yml ./config/

EXPOSE 8000
CMD ["./app"]