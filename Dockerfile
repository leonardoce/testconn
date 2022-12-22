# Step 1: build image
FROM golang:1.19 AS builder

# Cache the dependencies
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

# Compile the application
COPY . /app
RUN ./scripts/build.sh

# Step 2: build the image to be actually run
FROM gcr.io/distroless/static
COPY --from=builder /app/bin/testconn /app/bin/testconn
ENTRYPOINT ["/app/bin/testconn"]
