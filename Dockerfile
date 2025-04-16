# Etapa de construção
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar os arquivos de módulo primeiro (para melhor cache)
COPY go.mod go.sum ./
# RUN go mod download

# Copiar o restante do código
COPY . .

# Construir a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

# Etapa de execução
FROM alpine:latest

WORKDIR /app

# Copiar o binário e assets necessários
COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/docs ./docs

# Variáveis de ambiente padrão
ENV PORT=3000

EXPOSE $PORT

# Comando de execução
CMD ["/app/api"]