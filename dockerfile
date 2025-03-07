FROM golang:1.23.5 AS builder

WORKDIR /app

COPY . /app

RUN go mod tidy

COPY air.toml /app/air.toml

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["air", "-c", "air.toml"]