FROM golang:alpine

ADD . /go/src/github.com/wlsc/api-lightweight-essential

RUN \
  cd /go/src/github.com/wlsc/api-lightweight-essential && \
  go get -v && \
  go build -o /srv/api-lightweight-essential && \
  rm -rf /go/src/*

EXPOSE 8888

WORKDIR /srv

CMD ["/srv/api-lightweight-essential"]