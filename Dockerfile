FROM golang:alpine as builder
RUN apk add --no-cache git mercurial
RUN go get -u github.com/golang/dep/...
RUN mkdir /go/src/app 
ADD . /go/src/app/
WORKDIR /go/src/app
RUN dep ensure
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /go/src/app/main /app/
WORKDIR /app
CMD ["./main"]