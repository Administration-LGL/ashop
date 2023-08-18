FROM golang:1.20 AS builder

ARG APP_RELATIVE_PATH

COPY . /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

ARG APP_RELATIVE_PATH

COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

# CMD ["./server", "-conf", "/data/conf"]
CMD [ "./server" ]