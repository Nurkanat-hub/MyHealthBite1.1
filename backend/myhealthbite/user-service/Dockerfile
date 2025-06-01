FROM golang:1.24.1-alpine

WORKDIR /app

# Копируем зависимости
COPY ../email-service ./email-service
COPY ../stats-service ./stats-service

# Копируем модульные файлы заранее
COPY user-service/go.mod user-service/go.sum ./user-service/

# Копируем весь сервис
COPY user-service/ ./user-service/

WORKDIR /app/user-service

RUN go mod tidy
RUN go build -o user-service ./cmd/main.go

EXPOSE 50051

CMD ["./user-service"]
