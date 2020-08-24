# Build image
FROM golang:1.14-buster as build
WORKDIR /go/src/app
COPY . .
RUN export GO111MODULE=on
RUN mkdir -p /go/src/gitlab.com && ln -s $PWD/vendor/gitlab.com/vdat /go/src/gitlab.com/vdat
RUN go mod download         ## Go 1.11+

RUN go build ./cmd/chatserver


# Target image
FROM scratch
WORKDIR /go/src/app
COPY --from=build /go/src/app/chatserver ./
COPY index.html ./
CMD ["./chatserver"]
