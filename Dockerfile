# BUILD
FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app .

# RUNTIME
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app /usr/local/bin/app

EXPOSE 50051

CMD ["app"]
