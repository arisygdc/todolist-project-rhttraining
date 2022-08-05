FROM golang:1.19.0-alpine3.15 AS builder

WORKDIR /app
COPY . .

RUN go build -o gateway main.go

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/gateway .

CMD ["./gateway"]