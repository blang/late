FROM golang:1.22.5-alpine3.20 as builder
ARG APP_VERSION
ENV APP_VERSION=$APP_VERSION
WORKDIR /go/src/late
COPY . /go/src/late/
RUN apk --no-cache add git ca-certificates zip
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN mkdir /tmp/build/
RUN mkdir /tmp/build/darwin_arm64 /tmp/build/darwin_amd64 /tmp/build/linux_amd64 /tmp/build/linux_arm64
RUN env GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/darwin_arm64/late ./cmd/late
RUN env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/darwin_amd64/late ./cmd/late
RUN env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/linux_amd64/late ./cmd/late
RUN env GOOS=linux GOARCH=arm64 go build -ldflags "-X main.SemVer=$APP_VERSION" -o /tmp/build/linux_arm64/late ./cmd/late
RUN cd /tmp/build/darwin_arm64 && zip ../late_darwin_arm64.zip late
RUN cd /tmp/build/darwin_amd64 && zip ../late_darwin_amd64.zip late
RUN cd /tmp/build/linux_amd64 && zip ../late_linux_amd64.zip late
RUN cd /tmp/build/linux_arm64 && zip ../late_linux_arm64.zip late

FROM scratch
COPY --from=builder /tmp/build/*.zip /release/
