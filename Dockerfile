FROM alpine:3.10.1
RUN mkdir -p chmod 666 /app
WORKDIR /app
COPY ./cmd/client/bin /app
ENTRYPOINT chmod +x client
