FROM golang:1.10 as builder
COPY . $GOPATH/src/github.com/shikloshi/worker-go/
WORKDIR $GOPATH/src/github.com/shikloshi/worker-go/

#RUN apk update && apk add git
RUN go get -d -v

RUN go build -o /go/bin/worker-go

FROM scratch

COPY --from=builder /go/bin/worker-go /go/bin/worker-go

#ENTRYPOINT ["/go/bin/worker-go"]
