FROM registry.redhat.io/rhel8/go-toolset AS builder

WORKDIR /app
COPY . .

RUN go build -o user-service main.go

FROM registry.redhat.io/ubi8/ubi

WORKDIR /app

COPY --from=builder /app/user-service .

CMD ["./user-service"]