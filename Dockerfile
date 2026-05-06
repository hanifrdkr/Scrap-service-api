FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine:3.19 AS runtime

RUN apk add --no-cache tzdata
ENV APP_COMMAND http

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["sh", "-c", "./app ${APP_COMMAND} ./config.yaml"]