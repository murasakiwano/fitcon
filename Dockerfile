# Use the offical Golang image to build the app: https://hub.docker.com/_/golang
FROM golang:1.21.3 as builder

# Copy code to the image
WORKDIR /go/src/github.com/murasakiwano/fitcon
COPY . .

# Build the app
RUN go build -v -o fitcon .

# Start a new image for production without build dependencies
FROM debian:bookworm-slim

# Copy the app binary from the builder to the production image
COPY --from=builder /go/src/github.com/murasakiwano/fitcon/fitcon /fitcon
COPY --from=builder /go/src/github.com/murasakiwano/fitcon/assets /assets

# Run the app when the vm starts
CMD ["/fitcon"]
