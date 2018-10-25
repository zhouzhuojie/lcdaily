FROM golang:1.10
WORKDIR /go/src/github.com/zhouzhuojie/lcdaily
ADD . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build
CMD ./lcdaily
