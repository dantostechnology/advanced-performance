FROM golang:alpine 

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GIN_MODE=release

# For email certificate
RUN apk add -U --no-cache ca-certificates

# Copy the code into the container
COPY media /media/

# Move to working directory /build
WORKDIR /build

# Copy the code from /app to the build folder into the container
COPY app .

# Configure the build (go.mod and go.sum are already copied with prior step)
RUN go mod download

# Build the application
RUN go build -o main .

WORKDIR /app

# Copy binary from build to main folder
RUN cp /build/main .

# For email certificate, the 2 lines
VOLUME /etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Export necessary port
EXPOSE 8050

# Command to run when starting the container
CMD ["/app/main"]