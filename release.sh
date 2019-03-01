#!/bin/bash
set -e -o pipefail

GIT_VERSION=$(git describe --tags --match "v[0-9]*" HEAD | grep "^v" | head -1)
if [ -n "$GIT_TAG" ]; then
   GIT_VERSION="v0.0.0"
fi
version="$GIT_VERSION"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo "$DIR"

docker build --build-arg HTTP_PROXY="${HTTP_PROXY}" --build-arg http_proxy="${http_proxy}" --build-arg HTTPS_PROXY="${HTTPS_PROXY}" --build-arg https_proxy="${https_proxy}" -t "build/late:${version}" --build-arg APP_VERSION="$version" -f "$DIR/Dockerfile" "$DIR"

echo "# Creating container to copy source from"
id=$(docker create "build/late:${version}" sh)
docker cp $id:'/release/late_darwin_amd64.zip' ./build/
docker cp $id:'/release/late_linux_amd64.zip' ./build/
echo "# Removing container"
docker rm -v $id
echo "# Success"
