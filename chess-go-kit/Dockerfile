FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .
#
# Build the application
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
COPY --from=0  /build/main .

# Export necessary port
EXPOSE 9000

COPY appconf/config.yaml /dist/appconf/config.yaml
# Command to run when starting the container
CMD ["/dist/main"]