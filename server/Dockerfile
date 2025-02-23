# Use official Go image to build the binary
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

# Use a lightweight image to run the app
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
