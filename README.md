# blinkt_go_examples

These examples go with the [Blinkt_go library](https://github.com/alexellis/blinkt_go/)

The [Pimoroni Blinkt!](https://shop.pimoroni.com/products/blinkt) is an 8x RGB LED add-on for the Raspberry Pi.

Running one of the examples:

```
export GOPATH=$HOME/go/
mkdir -p /home/pi/go/src/github.com/alexellis/
git clone https://github.com/alexellis/blinkt_go_examples

cd progress
go get
sudo -E go run app.go
```

The Pimoroni Python examples are available here:

https://github.com/pimoroni/blinkt/tree/master/examples

PRs are welcome to port examples across to Go.

![](https://cdn.shopify.com/s/files/1/0174/1800/products/Blinkt_-_Zero-2_1024x1024.JPG?v=1466525645)
