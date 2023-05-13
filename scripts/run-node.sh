#!/bin/sh

set -eo pipefail

docker buildx create --name core-builder || true
docker buildx use core-builder
docker buildx build \
    --build-arg BUILDPLATFORM=linux/amd64 \
    --build-arg GOOS=linux \
    --build-arg GOARCH=amd64 \
    -t core:local-amd64 \
    --load \
    -f Dockerfile .

DOCKER_BUILDKIT=0 docker build -t doxchain:arch -f tools/env/Dockerfile.arch --platform linux/amd64 tools/env

docker compose -f tools/env/docker-compose.yml up -d