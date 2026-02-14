# =============================================
# Stage 1: Build binary Go
# =============================================
FROM golang:1.25-alpine AS builder

# Install gcc untuk CGO (diperlukan oleh go-sqlite3)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Salin dependency files dulu untuk cache layer
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh source code
COPY . .

# Build binary dengan CGO enabled (untuk sqlite3)
RUN CGO_ENABLED=1 GOOS=linux go build -o portfolio-server ./cmd/server

# =============================================
# Stage 2: Runtime image minimal
# =============================================
FROM alpine:latest

# Install ca-certificates untuk HTTPS dan timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Salin binary dari builder stage
COPY --from=builder /app/portfolio-server .

# Salin file statis dan template
COPY --from=builder /app/web ./web
COPY --from=builder /app/migrations ./migrations

# Buat direktori untuk database
RUN mkdir -p /app/data

# Expose port
EXPOSE 8080

# Jalankan server
CMD ["./portfolio-server"]
