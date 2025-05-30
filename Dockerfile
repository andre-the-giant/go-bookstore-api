FROM golang:1.24 AS builder

WORKDIR /app

# Clone from GitHub
RUN apt-get update && apt-get install -y git \
  && git clone https://github.com/andre-the-giant/go-bookstore-api.git . \
  && go mod tidy \
  && go build -o server .

# Final image
FROM golang:1.24

WORKDIR /app

# Copy compiled binary and init.sql from builder
COPY --from=builder /app/server .
COPY --from=builder /app/init.sql .

CMD ["./server"]
