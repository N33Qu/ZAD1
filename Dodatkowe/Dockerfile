# Stage 1: Budowanie aplikacji
FROM golang:1.17 AS builder

# Ustawienie środowiska
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# Skopiowanie plików aplikacji
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY .. .

# Kompilacja aplikacji
RUN go build -o server .

# Budowanie projektu React
FROM node:14 AS frontend-builder

WORKDIR /app

RUN npx create-react-app frontend

WORKDIR /app/frontend
# Skopiuj resztę plików frontendowych
COPY frontend/src ./src

# Zbuduj aplikację frontendową
RUN npm run build

# Obraz końcowy
FROM alpine:latest

WORKDIR /app

# Skopiuj pliki z budowania Go
COPY --from=builder /app/server /server

# Skopiuj zbudowany kod frontendowy
COPY --from=frontend-builder /app/frontend/build ./frontend

COPY index.html /app/frontend

# Uruchom aplikację
CMD ["/server"]
