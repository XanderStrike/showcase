#!/bin/bash

# update dependencies
docker run --rm -it -v $PWD:/go/src/github.com/xanderstrike/showcase -w /go/src/github.com/xanderstrike/showcase treeder/glide update

# build binary
docker run --rm -v "$PWD":/go/src/github.com/xanderstrike/showcase -w /go/src/github.com/xanderstrike/showcase iron/go:dev go build -o showcase-docker

# build docker image
docker build -t xanderstrike/showcase .

rm showcase-docker

read -p "Push to dockerhub? " -n 1 -r
echo    # (optional) move to a new line
if [[ $REPLY =~ ^[Yy]$ ]]
then
  docker push xanderstrike/showcase:latest
fi
