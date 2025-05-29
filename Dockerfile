# Build stage
FROM golang:1.24-alpine as builder

WORKDIR /app
COPY . .

RUN go build -o pod-kicker .

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from builder
COPY --from=builder /app/pod-kicker .

# Make sure it’s executable
RUN chmod +x pod-kicker

# Install kubectl
RUN apk add --no-cache curl bash && \
    curl -LO https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    rm -f kubectl

# Run the pod-kicker binary
CMD ["./pod-kicker"]
