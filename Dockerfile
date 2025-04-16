FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/docs ./docs

# Variáveis de ambiente padrão
ENV PORT=3000

EXPOSE $PORT

# Comando de execução
CMD ["/app/api"]