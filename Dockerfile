FROM golang:latest as builder
RUN mkdir -p $GOPATH/src/deployer
WORKDIR $GOPATH/src/deployer
ADD . .
RUN CGO_ENABLED=0 go install deployer && mv $GOPATH/bin/deployer /bin/deployer

FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /bin/
COPY --from=builder /bin/deployer .