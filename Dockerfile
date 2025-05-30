FROM golang:1.24

WORKDIR /app

# Clone your public repo
RUN apt-get update && apt-get install -y git \
  && git clone https://github.com/andre-the-giant/go-bookstore-api.git . \
  && go mod tidy


RUN go build -o server .

# Run the app
CMD ["./server"]
