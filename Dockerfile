# --- STAGE 1: Build the Frontend ---
FROM node:20-alpine AS frontend-builder
WORKDIR /app/web
COPY web/package*.json ./
RUN npm install
COPY web/ ./
RUN npm run build

# --- STAGE 2: Build the Go Backend ---
FROM golang:1.26-alpine AS backend-builder
WORKDIR /app
# Install git for fetching dependencies if needed
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy the built frontend from Stage 1 into the Go 'web' folder
# This allows Go's 'embed' package to see the built files
COPY --from=frontend-builder /app/web/dist ./web/dist
# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

# --- STAGE 3: Final Production Image ---
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy only the compiled binary from Stage 2
COPY --from=backend-builder /app/server .
# Copy .env.example as a reference (optional)
COPY --from=backend-builder /app/.env.example .

# Expose the port your app runs on
EXPOSE 8080

# Run the server
CMD ["./server"]