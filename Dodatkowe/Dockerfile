# Stage 1: Budowanie aplikacji
FROM golang:latest AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
    
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY .. .

RUN go build -o server .

# Budowanie projektu React
FROM node:14 AS frontend-builder

WORKDIR /app

RUN npx create-react-app frontend

WORKDIR /app/frontend

COPY frontend/src ./src

RUN npm run build

# Obraz końcowy
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server /server

COPY --from=frontend-builder /app/frontend/build ./frontend

COPY index.html /app/frontend

CMD ["/server"]
