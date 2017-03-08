FROM resin/rpi-raspbian

# Install Go 1.7.5 and enough of a build-chain to build wiringpi (in C)
RUN apt-get update && \
    apt-get install -qy build-essential wiringpi git curl ca-certificates && \
    curl -sSLO https://storage.googleapis.com/golang/go1.7.5.linux-armv6l.tar.gz && \
	mkdir -p /usr/local/go && \
	tar -xvf go1.7.5.linux-armv6l.tar.gz -C /usr/local/go/ --strip-components=1

ENV PATH=$PATH:/usr/local/go/bin/
ENV GOPATH=/go/

RUN mkdir -p /go/src/github.com/alexellis/blinkt_go_examples/resistor_clock
WORKDIR /go/src/github.com/alexellis/blinkt_go_examples/resistor_clock

COPY app.go	.
RUN go get -d -v

RUN go build

CMD ["./resistor_clock"]
