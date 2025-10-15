FROM golang:latest
LABEL authors="chunan"

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 6969

ENV GOFLAGS="-buildvcs=false"

CMD ["air", "-c", ".air.toml"]
