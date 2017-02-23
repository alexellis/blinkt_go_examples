### Synchronised Blinkt via Redis

We have sender code, receiver code and then Redis as the connective tissue using a channel to exchange information.


* Run this on your RPis that have a Blinkt attached:

```
$ docker run -v /sys:/sys -e ADDR=192.168.2.21 -d alexellis2/redis-view
```

*Alter the ADDR to match the IP of the machine with Redis*

* To run Redis on an RPi:

```
$ docker run --name redis -d -p 6379:6379 alexellis2/redis-armhf
```

If you want to debug messages on Redis:

```
$ docker run -ti redis-armhf redis-cli -h 192.168.2.21
```

* Run this on your "sender"

```
$ cd redis_view
$ go run app.go
```

#### Docker Swarm

For Swarm do the following:

```
$ docker network create --driver overlay --attachable=true blinkt

$ docker service create --name redis --replicas=1 --network=blinkt --publish 6379:6379 \
  alexellis2/redis-armhf:latest

$ docker service create --name redis_view --mount type=bind,source=/sys,destination=/sys \
  -e ADDR=redis --network=blinkt \
  --mode=global alexellis2/redis-view:latest

# optionally: --constraint='node.role != manager' 
```

You can invoke the sender as a service like this:

```
$ docker service rm redis_sender; docker service create --name redis_sender --replicas=1 --network=blinkt  -e ADDR=redis  alexellis2/redis-view:latest
```

