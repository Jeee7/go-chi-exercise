# Stage 1: Build
FROM golang:1.21-alpine AS builder

# Set working dir ke /app
WORKDIR /app

# Copy go.mod dan go.sum dari root project (di atas server)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy semua source code ke /app (seluruh project)
COPY . .

# Build binary dari main.go di folder server
RUN CGO_ENABLED=0 GOOS=linux go build -o backend ./server

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copy binary hasil build dari builder
COPY --from=builder /app/backend .

# Expose port yang kamu pakai (sesuaikan dengan port server kamu)
EXPOSE 8080

# Jalankan binary
CMD ["./backend"]
