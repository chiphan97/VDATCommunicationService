# Build image
FROM golang:1.14-alpine as build
WORKDIR /go/src/app

COPY go.* modules/ ./
RUN go mod download
COPY . .
RUN mkdir -p /go/src/gitlab.com && ln -s $PWD/modules/gitlab.com/vdat /go/src/gitlab.com/vdat

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/chatserver ./cmd/chatserver

## Build migration
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/migrator ./migration

## BUILD ANGULAR WEBAPP
FROM node:12-alpine AS angular-build
WORKDIR /usr/src/app
COPY website/package.json ./
RUN npm install
COPY ./website .
#RUN npm run build:prod
RUN npm run build:staging

# Target image
# FROM gcr.io/distroless/base-debian10
# FROM ubuntu:20.04
FROM alpine:latest
WORKDIR /app
RUN apk add ca-certificates
COPY --from=build /go/bin/chatserver ./
COPY --from=build /go/bin/migrator ./
COPY --from=build /go/src/app/migration/ ./migration/
COPY --from=angular-build /usr/src/app/dist ./public 

EXPOSE 5000
CMD ["/app/migrator", "&&", "/app/cmd/chatserver"]
