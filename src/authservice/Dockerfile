FROM golang:1.19.0-alpine3.15 AS builder

WORKDIR /app
COPY . .

RUN apk add build-base
RUN go build -o auth-service main.go

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/auth-service .

CMD ["./auth-service"]