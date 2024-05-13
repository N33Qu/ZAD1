# Stage 1: Budowanie aplikacji
FROM golang:latest AS builder

WORKDIR /app

RUN mkdir -p /clonedRepo
RUN --mount=type=secret,id=ghlab_access \
    mkdir -p ~/.ssh && \
    cp /run/secrets/ghlab_access ~/.ssh/id_rsa && \
    chmod 600 ~/.ssh/id_rsa && \
    ssh-keyscan github.com >> ~/.ssh/known_hosts &&  \
    git clone git@github.com:N33Qu/ZAD1.git clonedRepo

RUN mv ./clonedRepo/Dodatkowe/go.mod .
RUN go mod download

RUN mv ./clonedRepo/Dodatkowe/* .

RUN go build -o server .

RUN ls -la

# Budowanie projektu React
FROM node:22.1-alpine3.18 AS frontend-builder

WORKDIR /app

RUN npx create-react-app frontend

WORKDIR /app/frontend

COPY --from=builder ../app/frontend/src ./src

RUN npm run build

# Obraz końcowy
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server ./server

COPY --from=frontend-builder /app/frontend/build ./frontend

COPY --from=builder app/index.html /app/frontend

RUN ls -la

CMD ["./server"]
