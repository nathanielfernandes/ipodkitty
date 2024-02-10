# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine3.19

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

EXPOSE 80
CMD ["./ipodkitty"]