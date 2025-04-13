# Estágio de build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/main .

# Estágio final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 3000
CMD ["/app/main"]