FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

ARG BIN

RUN mkdir -p /shaouai/bin
COPY ${BIN} /shaouai/bin/${BIN}
