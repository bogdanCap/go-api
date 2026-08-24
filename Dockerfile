# Use the official Golang image for building
FROM golang:1.26.0 AS base
# Set working directory
WORKDIR /root
# Copy Go modules and dependencies
COPY go.mod ./
COPY go.sum ./
#11
#COPY vendor ./vendor


# or - RUN go mod vendor or RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-s -w" -o myapp .
##11 this do not need if i use vendor
RUN go mod download

# ==============================================================================
# STAGE 2: Local Live-Development (Runs Air)
# ==============================================================================
FROM base AS development
# Install Air for hot-reloading
# in my case test or ://github.com
#RUN go install test
# Copy the pre-downloaded binary directly from your computer into the container path
#COPY air /usr/local/bin/air
#RUN chmod +x /usr/local/bin/air
#ENV PATH="/go/bin:${PATH}"
RUN go install github.com/cosmtrek/air@v1.40.4

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Install Air directly into /usr/local/bin
######RUN GOBIN=/usr/local/bin go install github.com/air-verse/air@latest

# Verify installation during image build
#RUN /usr/local/bin/air -v

# Mount points happen via docker-compose; copy code for fallback
COPY . .
#RUN mkdir -p /app/tmp
RUN mkdir -p /root/tmp && chmod -R 777 /root/tmp

#RUN go build -o myapp ./cmd

EXPOSE 8080
# Run Air inside the container
##11 go build -buildvcs=false -o ./tmp/main ./cmd
CMD ["/go/bin/air", "-c", ".air.toml"]


# ==============================================================================
# STAGE 3: Production Builder
# ==============================================================================
FROM base AS builder
# Copy source code
COPY . .
#RUN go build -o ./tmp/main ./cmd
##11 RUN go build -mod=vendor -o app ./cmd
RUN go build -o app ./cmd

# ==============================================================================
# STAGE 4: Secure Production Image
# ==============================================================================
FROM alpine:3.21 AS production
# Build the application
#RUN go build -o main .
# Use a minimal base image for final deployment
#FROM alpine:latest
# Set working directory in the container
WORKDIR /root
#old version - WORKDIR /
RUN apk --no-cache add ca-certificates
# Copy ONLY the final binary from the builder stage
#COPY --from=builder /app/main app/tmp/main
COPY --from=builder /root/app root/app

# Copy the built binary from the builder stage
#COPY --from=builder /app/main .
# Expose the application port
EXPOSE 8080
# Run the application
CMD ["/root/app"]