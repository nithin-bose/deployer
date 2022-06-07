FROM golang:1.18 as builder
ADD . .
ARG DEPLOYER_VERSION
RUN CGO_ENABLED=0 go install -ldflags "-X deployer/pkg.Version=${DEPLOYER_VERSION}" deployer && mv $GOPATH/bin/deployer /bin/deployer

FROM alpine
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
WORKDIR /bin/
COPY --from=builder /bin/deployer .
