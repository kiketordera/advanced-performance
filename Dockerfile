FROM golang:alpine as builder

RUN apk update && apk add git && apk add ca-certificates 
# For email certificate
RUN apk add -U --no-cache ca-certificates

COPY . $GOPATH/src/github.com/kiketordera/advanced-performance/
WORKDIR $GOPATH/src/github.com/kiketordera/advanced-performance/

RUN go get -d -v $GOPATH/src/github.com/kiketordera/advanced-performance

# For Cloud Server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/advanced-performance $GOPATH/src/github.com/kiketordera/advanced-performance

FROM scratch
COPY --from=builder /go/bin/advanced-performance /advanced-performance
COPY --from=builder /go/src/github.com/kiketordera/advanced-performance/media/ /media/
COPY --from=builder /go/src/github.com/kiketordera/advanced-performance/*.html /
# For email certificate
VOLUME /etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8050/tcp

ENV GOPATH /go
ENTRYPOINT ["/advanced-performance"]
