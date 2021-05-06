FROM alpine:3.10.1
RUN mkdir -p /app
WORKDIR /app
COPY ./cmd/client/bin /app
ENTRYPOINT chmod 666 ["./client"]
