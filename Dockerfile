FROM golang:alpine3.19 as builder

WORKDIR /app

COPY src/presentation/github .

RUN go mod tidy
RUN go build -o app.golang

FROM alpine:latest

COPY --from=builder /app/app.golang /app/app.golang

WORKDIR /app

ENTRYPOINT [ "/app/app.golang" ]