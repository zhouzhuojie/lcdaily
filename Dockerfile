FROM golang:1.10 as builder
WORKDIR /go/src/github.com/zhouzhuojie/lcdaily
ADD . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

FROM alpine:3.6
RUN apk add --no-cache libc6-compat ca-certificates curl
WORKDIR /go/src/github.com/zhouzhuojie/lcdaily
COPY --from=builder /go/src/github.com/zhouzhuojie/lcdaily/lcdaily ./lcdaily
COPY --from=builder /go/src/github.com/zhouzhuojie/lcdaily/email_template.html ./email_template.html
CMD ./lcdaily
