# dev stage
FROM golang:latest AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]

# build stage
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/api

# runtime stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

CMD ["/app/main"]
