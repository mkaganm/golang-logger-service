FROM golang:1.20.5-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o logger_service ./cmd/main.go

RUN chmod +x /app/logger_service

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/logger_service /app
COPY ./app.env ./app.env

EXPOSE 3003:3003

CMD ["/app/logger_service"]