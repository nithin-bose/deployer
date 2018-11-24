FROM golang:latest as builder
RUN mkdir -p $GOPATH/src/deployer
WORKDIR $GOPATH/src/deployer
ADD . .
RUN CGO_ENABLED=0 go install deployer && mv $GOPATH/bin/deployer /bin/deployer

FROM alpine

RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache bash
RUN apk add --no-cache curl
RUN apk add --no-cache openssl

COPY install-helm.sh .
RUN cat install-helm.sh | bash

WORKDIR /bin/
COPY --from=builder /bin/deployer .

WORKDIR /charts/apps