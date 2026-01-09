# ---------- BUILD ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código
COPY . .

# Build do binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o api ./cmd/api

# ---------- RUNTIME ----------
FROM alpine:latest

WORKDIR /app

# Certificados (HTTPS, JWT, etc)
RUN apk --no-cache add ca-certificates

# Copia o binário
COPY --from=builder /app/api .

# Copia o .env (opcional, pode usar env do compose)
COPY .env .env

EXPOSE 8080

CMD ["./api"]
