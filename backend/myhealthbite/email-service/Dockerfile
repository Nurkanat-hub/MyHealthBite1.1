FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o email-service ./cmd/main.go

EXPOSE 50058

CMD ["./email-service"]
