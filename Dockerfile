FROM golang:1.24-alpine AS builder

# Adicionar dependências essenciais para build
RUN apk add --no-cache git

WORKDIR /app

# Copiar apenas os arquivos necessários primeiro
COPY go.mod go.sum ./
RUN go mod download

# Depois copiar o resto do código
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/api ./cmd/api

FROM alpine:latest

# Adicionar certificados CA e criar usuário não-root
RUN apk add --no-cache ca-certificates && \
    adduser -D appuser

WORKDIR /app

COPY --from=builder /app/bin/api /app/api
COPY --from=builder /app/docs ./docs

# Mudar proprietário dos arquivos
RUN chown -R appuser:appuser /app

# Mudar para usuário não-root
USER appuser

# Variáveis de ambiente padrão
ENV PORT=3000

EXPOSE $PORT

# Comando de execução
CMD ["/app/api"]