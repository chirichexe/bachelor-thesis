# Usa un'immagine Go per compilare il binario
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN go build -o device-agent main.go

# Immagine finale
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/device-agent .

ENTRYPOINT ["./device-agent"]
