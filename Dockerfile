FROM golang:1.17 as builder

ENV GOPATH=""
ENV GOMODULE="on"

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD app app
ADD cmd cmd
ADD x x

COPY Makefile .
RUN make build

FROM ubuntu:20.04

COPY --from=builder /go/build/icad /bin/icad

ENTRYPOINT ["icad"]
