FROM golang:1.23-alpine

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY . .

RUN go mod tidy && go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

EXPOSE 3000

# Commande pour lancer l'application
CMD ["./main"]