FROM golang:1.22.5-alpine3.20 as builder
ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION
WORKDIR /go/src/late
COPY . /go/src/late/
RUN apk --no-cache add git ca-certificates zip
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN mkdir /tmp/build/
RUN env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/late_darwin_amd64 ./cmd/late
RUN env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/late_linux_amd64 ./cmd/late
RUN cd /tmp/build && zip late_darwin_amd64.zip late_darwin_amd64
RUN cd /tmp/build && zip late_linux_amd64.zip late_linux_amd64

FROM scratch
COPY --from=builder /tmp/build/*.zip /release/
