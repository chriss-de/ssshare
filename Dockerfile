FROM docker.io/library/golang:1.24-alpine as golang

RUN set -ex; \
    apk add --no-cache \
      make ; \
    update-ca-certificates ; \
    rm -fr /var/cache/apk ;

WORKDIR /BUILD/

ENV CGO_ENABLED=0
ADD . .
RUN make all NAME=ssshare



FROM docker.io/library/alpine:3.21
COPY --from=golang /BUILD/out/bin/ssshare /bin
