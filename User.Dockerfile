FROM golang:1.17 as builder
ENV CGO_ENABLED 0

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -v -o user-service cmd/user/main.go

FROM alpine:3.15

RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/config/config.prod.yml /app/config/config.prod.yml
COPY --from=builder /app/user-service /app/user-service
COPY --from=builder /app/keys /app/keys

WORKDIR /app

# Run the web service on container startup.
CMD ["./user-service"]
