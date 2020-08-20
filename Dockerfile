# Build image
FROM golang:1.14-buster as build
WORKDIR /go/src/app
COPY go.* ./
RUN go mod download         ## Go 1.11+
COPY . .
RUN go build ./cmd/chatserver


# Target image
FROM golang:1.14-buster
COPY --from=build /go/src/app/chatserver /
CMD ["/chatserver"]
