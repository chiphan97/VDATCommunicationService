FROM golang:1.14-alpine

WORKDIR /go/src/app

ENV ENV_MODE prod

COPY ./chat-service .

CMD ["./chat-service"]