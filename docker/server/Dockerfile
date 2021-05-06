FROM alpine:3.10.1
RUN mkdir -p /app
WORKDIR /app
COPY ./cmd/server/bin /app
ENTRYPOINT chmod 666 ["./backend"]
