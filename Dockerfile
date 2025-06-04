# syntax=docker/dockerfile:1.4

FROM golang:1.24.3-alpine3.21

WORKDIR /app

# Install air for live reload
RUN go install github.com/air-verse/air@latest
RUN go install github.com/google/wire/cmd/wire@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air"]
