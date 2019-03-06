#!/bin/sh -x

mkdir -p "$PWD/build"

export PROVIDER="terraform-provider-$1"
export CGO_ENABLED=0

build_for_arch() {
 export GOOS="$1"
 export GOARCH="$2"

 go build -o "$PWD/build/$PROVIDER-$GOOS-$GOARCH" "github.com/phogolabs/terraform-provider/cmd/$PROVIDER"
}

build_for_arch "linux" "amd64"
build_for_arch "linux" "386"

build_for_arch "darwin" "amd64"
build_for_arch "darwin" "386"
