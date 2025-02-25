FROM golang:1.24.0 as builder

# Set destination for COPY
WORKDIR /build
COPY .env .

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
RUN go build -o /main src/main.go

FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]

# Run
CMD ["/simple-API-dnd"]