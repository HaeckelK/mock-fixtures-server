# Use pre built fixtures server as builder
FROM docker pull haeckelk/fixtures-server:latest as builder

# Copy in own fixtures
COPY ./fixtures /fixtures

# Copy in own router.go to overwrite default
COPY ./router.go /app/router.go

RUN go build

# Build fresh image which contains go binaries and fixtures
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/mock-fixtures-server ./mock-fixtures-server
COPY ./fixtures /fixtures
CMD ["./mock-fixtures-server"]  