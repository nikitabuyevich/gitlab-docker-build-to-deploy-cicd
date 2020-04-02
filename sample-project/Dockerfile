############################
# STEP 1 build executable binary
############################
# golang debian buster 1.14.1 linux/amd64
# https://github.com/docker-library/golang/blob/master/1.14/buster/Dockerfile
FROM golang@sha256:eee8c0a92bc950ecb20d2dffe46546da12147e3146f1b4ed55072c10cacf4f4c as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

# Set the working directory and copy its contents
WORKDIR $GOPATH/src/app/
COPY . .

# Download and verify dependencies
RUN go mod download
RUN go mod verify

# Disable cross-platform compilation and build the static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/app .

# Build tests static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go test \
      -ldflags='-w -s -extldflags "-static"' -a \
      -c $GOPATH/src/app \
      -o /go/bin/tests

############################
# STEP 2 build a small image
############################
# Using static nonroot image
# User:group is nobody:nobody, uid:gid = 65534:65534
# Get the sha256 from the "latest" tagged image
FROM gcr.io/distroless/static@sha256:c6d5981545ce1406d33e61434c61e9452dad93ecd8397c41e89036ef977a88f4

# Copy our static executables
COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /go/bin/tests /go/bin/tests

# Run the app
ENTRYPOINT ["/go/bin/app"]
