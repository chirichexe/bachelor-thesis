FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o device-agent main.go

# Immagine finale
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/device-agent .

EXPOSE 8080

ENTRYPOINT ["./device-agent"]
