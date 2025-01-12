# Stage 1: Build the Go application
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /app/bin/qryptic-gateway /app/cmd/gateway/main.go

# Stage 2: Create the final image
FROM alpine:latest
RUN apk --no-cache add wireguard-tools iptables bash
COPY --from=builder /app/bin/qryptic-gateway /usr/local/bin/qryptic-gateway
CMD ["/usr/local/bin/qryptic-gateway"]