FROM alpine:3.8

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/ustrugany/classifier

ADD classifier /go/src/github.com/ustrugany/classifier
COPY config config
COPY web web

CMD ["./classifier", "server"]
