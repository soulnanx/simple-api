FROM golang:1.21-alpine AS builder 

WORKDIR /app

# Primeiro copia apenas os arquivos de módulo
COPY go.mod go.sum ./

# Baixa as dependências
RUN go mod download

# Copia o resto do código
COPY . .

# Constrói a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 3000
CMD ["/app/main"]