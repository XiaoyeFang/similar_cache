FROM golang:latest AS build-env
ADD . $GOPATH/src/similar_cache
WORKDIR $GOPATH/src/similar_cache

ENV CGO_ENABLED=0
RUN mkdir -p /output && go build -o /output/similar_cache

FROM alpine:latest
MAINTAINER lizhiyuan "lizhiyuan@apkpure.net"
WORKDIR /app
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/*
COPY --from=build-env /output/similar_cache /app/
CMD ["/app/similar_cache"]