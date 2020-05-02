FROM golang:alpine

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux
#   \ GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy download update dependencies
RUN go get github.com/golang/dep/cmd/dep
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export port
EXPOSE 8080

# Start the container
CMD ["/dist/main"]