FROM alpine

ENV ENV_MODE prod

COPY ./chat-service .

CMD ["./chat-service"]