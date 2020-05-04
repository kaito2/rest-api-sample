FROM golang:1.13 as base

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go get goa.design/goa/v3@v3.0.6
RUN go get goa.design/goa/v3/...@v3.0.6

COPY go.* ./
RUN go mod download

COPY . .
RUN ./scripts/goa_gen.sh

# it will take the flags from the environment
RUN go build -o main ./cmd/echo_server

### Certs
FROM alpine:latest as certs
RUN apk --update add ca-certificates

### App
FROM alpine:latest as app
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=base /app/main /
CMD ["/main"]