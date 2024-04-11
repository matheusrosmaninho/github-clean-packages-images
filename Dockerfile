FROM golang:alpine3.19 as builder

WORKDIR /app

COPY src .

WORKDIR /app/presentation/local

RUN go mod tidy
RUN go build -o app.golang

FROM alpine:latest

COPY --from=builder /app/presentation/local/app.golang /app/app.golang
COPY --from=builder /app/presentation/local/.env /app/.env

WORKDIR /app

ENTRYPOINT [ "/app/app.golang" ]