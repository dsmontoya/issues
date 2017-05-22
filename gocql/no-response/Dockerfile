FROM golang:1.8.0
ADD . /go/src/github.com/dsmontoya/issues/gocql/no-response
WORKDIR /go/src/github.com/dsmontoya/issues/gocql/no-response
RUN apt-get install git gcc g++
RUN go get
RUN go install -race github.com/dsmontoya/issues/gocql/no-response
CMD no-response
