FROM alexellis2/go-armhf:1.9 as build

# Official image only works on RPi2/3 due to Docker issue with official images. 
# FROM golang:1.9.2 as build
RUN apk add --no-cache git

ENV GOPATH=/go
RUN go get -d -u github.com/alexellis/blinkt_go/sysfs/gpio

RUN mkdir -p /go/src/github.com/alexellis/blink_go_examples/example
WORKDIR /go/src/github.com/alexellis/blinkt_go_examples/example

COPY . .

RUN GOARM=6 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/bin/app .

FROM scratch

COPY --from=build /usr/bin/app /app

CMD ["/app"]
