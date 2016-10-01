FROM golang:1.7

RUN apt-get update && apt-get install -y wget
RUN wget https://github.com/jwilder/dockerize/releases/download/v0.2.0/dockerize-linux-amd64-v0.2.0.tar.gz
RUN tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.2.0.tar.gz

RUN mkdir -p /go/src/github.com/porthos-rpc/porthos-dashboard
WORKDIR /go/src/github.com/porthos-rpc/porthos-dashboard

RUN go get -u github.com/kardianos/govendor
