#!/usr/bin/env bash

set -x
set -e
set -o pipefail

if [ ! -d "/root/.docker" ]; then
    mkdir -p /root/.docker
fi

mv docker-config/docker-ecr-config.json /root/.docker/config.json
