FROM golang
MAINTAINER Alex Standke "xanderstrike@gmail.com"

ADD . /go/src/github.com/XanderStrike/showcase

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

VOLUME ["/watch"]

EXPOSE 8080
ENTRYPOINT revel run github.com/XanderStrike/showcase prod 8080

