# blinkt_go_examples

These examples go with the [Blinkt_go library](https://github.com/alexellis/blinkt_go/)

The [Pimoroni Blinkt!](https://shop.pimoroni.com/) is an 8x RGB LED add-on for the Raspberry Pi.

Running one of the examples:

```
export GOPATH=$HOME/go/
mkdir -p /home/pi/go/src/github.com/alexellis/
git clone https://github.com/alexellis/go_blinkt_examples

cd progress
go get
sudo -E go run app.go
```

The Pimoroni Python examples are available here:

https://github.com/pimoroni/blinkt/tree/master/examples

PRs are welcome to port examples across to Go.
