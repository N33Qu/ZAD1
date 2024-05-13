# Stage 1: Budowanie aplikacji
FROM golang:alpine AS builder

# Ustawienie środowiska
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Skopiowanie plików aplikacji
WORKDIR /app
COPY . .

# Kompilacja aplikacji
RUN go build -o server .

# Stage 2: Tworzenie obrazu kontenera
FROM scratch

# Kopiowanie binarnej aplikacji z poprzedniego etapu
COPY --from=builder /app/server /server

# Uruchomienie serwera HTTP
CMD ["/server"]
