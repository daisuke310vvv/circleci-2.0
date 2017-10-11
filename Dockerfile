FROM golang:1.8.1

EXPOSE 9090

ADD workspace/server /go/

CMD ["/go/server"]
