FROM golang:1.25.0-alpine AS builder

WORKDIR /app

# Copia arquivos de dependência primeiro (melhor cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código
COPY . .

# Compila binário estático
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api

# Runtime stage 
FROM alpine:latest

WORKDIR /app

# Instala certificados SSL
RUN apk --no-cache add ca-certificates

# Copia binário do stage anterior
COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]