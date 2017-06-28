FROM golang:1.8

RUN curl https://glide.sh/get | sh

COPY . /go/src/github.com/crucial7/api-geo
WORKDIR /go/src/github.com/crucial7/api-geo

RUN chmod +x build.sh
RUN ./build.sh

CMD ["./index"]
