# Build image
FROM golang:1.14-alpine as build
WORKDIR /go/src/app

EXPOSE 5000

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

COPY . .
RUN mkdir -p /go/src/gitlab.com && ln -s $PWD/modules/gitlab.com/vdat /go/src/gitlab.com/vdat
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/chatserver ./cmd/chatserver

## Build migration
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
          -ldflags='-w -s -extldflags "-static"' -a \
          -o /go/bin/migration ./migration

## BUILD ANGULAR WEBAPP
FROM node:12-alpine AS angular-build
WORKDIR /usr/src/app
COPY website/package.json ./
RUN npm install
COPY ./website .
RUN npm run build:prod

# Target image
FROM gcr.io/distroless/base-debian10
WORKDIR /go/src/app
COPY --from=build /go/bin/chatserver ./
COPY --from=build /go/bin/migration ./
COPY --from=build /go/src/app/migration/*.sql ./migration/
COPY --from=angular-build /usr/src/app/dist ./public
CMD ["./migration", "&&", "./chatserver"]
