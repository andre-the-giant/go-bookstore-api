FROM golang:1.24 AS builder

# WORKDIR can be anything
WORKDIR /go/src/app

# Clone from your public GitHub repo
RUN apt-get update && apt-get install -y git \
  && git clone https://github.com/andre-the-giant/go-bookstore-api.git /repo \
  && ls -l /repo \
  && cd /repo \
  && go mod tidy \
  && go build -o /go/bin/server /repo/main.go

# Final image
FROM golang:1.24

WORKDIR /app

# Copy the binary and the init.sql
COPY --from=builder /go/bin/server .
COPY --from=builder /repo/init.sql .

CMD ["./server"]
