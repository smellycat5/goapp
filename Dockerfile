FROM golang:latest
LABEL authors="chunan"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 6969
