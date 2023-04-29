FROM golang:1.20 as builder
RUN mkdir -p /src/deployer
WORKDIR /src/deployer

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

ADD . .
ARG DEPLOYER_VERSION
RUN CGO_ENABLED=0 go install -ldflags "-X deployer/pkg.Version=${DEPLOYER_VERSION}" deployer && mv $GOPATH/bin/deployer /bin/deployer

FROM alpine
RUN apk add --no-cache git
RUN apk add --no-cache ca-certificates
WORKDIR /bin/
COPY --from=builder /bin/deployer .
