FROM golang:1.8

COPY . /go/src/github.com/crucial7/api-geo
WORKDIR /go/src/github.com/crucial7/api-geo

CMD ["./run.sh"]
