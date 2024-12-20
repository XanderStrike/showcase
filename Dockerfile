FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o showcase .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/showcase .
COPY index.html .
COPY intro.html .

RUN adduser -D appuser && \
    mkdir config && \
    chown -R appuser:appuser /app

VOLUME /app/config

USER appuser

EXPOSE 8080
CMD ["./showcase"]
