# Build image
FROM golang:1.14-alpine as build
WORKDIR /go/src/app

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

COPY . .
RUN mkdir -p /go/src/gitlab.com && ln -s $PWD/modules/gitlab.com/vdat /go/src/gitlab.com/vdat
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/chatserver ./cmd/chatserver


# Target image
FROM scratch
WORKDIR /go/src/app
COPY index.html ./
COPY --from=build /go/bin/chatserver ./
CMD ["./chatserver"]
