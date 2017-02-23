#!/bin/sh

docker build --build-arg http_proxy=$http_proxy --build-arg https_proxy=$https_proxy -t alexellis2/redis-sender:builder . -f Dockerfile.build && \
docker create --name sysfs-builder alexellis2/redis-sender:builder && \
docker cp sysfs-builder:/go/src/github.com/alexellis/blinkt_go_examples/sysfs/progress/app . && \
docker rm -f sysfs-builder && \
docker build -t alexellis2/redis-sender:latest .



