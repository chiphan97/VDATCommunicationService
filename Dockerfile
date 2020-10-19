# Build image
FROM golang:1.14-alpine as build
WORKDIR /go/src/app

EXPOSE 5000

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

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
RUN npm run build:prod

# Target image
#FROM gcr.io/distroless/base-debian10
FROM ubuntu:20.04
WORKDIR /app
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=build /go/bin/chatserver ./
COPY --from=build /go/bin/migrator ./
COPY --from=build /go/src/app/migration/ ./migration/
COPY --from=angular-build /usr/src/app/dist ./public 
RUN ls -la
CMD [ "sh", "-c", "./migrator && echo 1 &&  ./chatserver"]
