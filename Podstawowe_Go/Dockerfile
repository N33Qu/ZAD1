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

#Doadanie etykiety z imieniem i nazwiskiem autora
LABEL author="Mikolaj Stasz"

#Przeprowadzenie healthcheck
HEALTHCHECK --interval=10s --timeout=1s \
 CMD curl -f http://localhost:8080/ || exit 1

# Uruchomienie serwera HTTP
CMD ["/server"]
