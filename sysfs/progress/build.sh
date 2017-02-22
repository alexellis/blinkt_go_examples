#!/bin/sh

docker build -t alexellis2/progress-blinkt:sysfs-builder .
docker create --name sysfs-builder alexellis2/progress-blinkt:sysfs-builder
docker cp sysfs-builder:/go/src/github.com/alexellis/blinkt_go_examples/sysfs/progress/progress .

docker rm -f sysfs-builder

docker build -t alexellis2/progress-blinkt:sysfs .

