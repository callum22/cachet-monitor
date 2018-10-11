FROM golang

RUN go get github.com/Sirupsen/logrus
RUN go get github.com/miekg/dns
ADD . /go/src/github.com/ArthurHlt/cachet-monitor
RUN go install github.com/ArthurHlt/cachet-monitor

ENTRYPOINT /go/bin/cachet-monitor
