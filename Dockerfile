# Build image
FROM golang:1.14-buster as build
WORKDIR /go/src/app
COPY go.* ./
RUN go mod download         ## Go 1.11+
COPY . .
RUN go build ./cmd/chatserver


# Target image
FROM scratch
WORKDIR /go/src/app
COPY --from=build /go/src/app/chatserver ./
COPY index.html ./
CMD ["./chatserver"]
