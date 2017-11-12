FROM golang:1.9.2 as build

ENV GOPATH=/go
RUN go get -d -u github.com/alexellis/blinkt_go/sysfs/gpio
RUN mkdir -p /go/src/github.com/alexellis/blinkt_go_examples/cheerlights
WORKDIR /go/src/github.com/alexellis/blinkt_go_examples/cheerlights
COPY . .
RUN GOARM=6 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/bin/app .

FROM scratch

COPY --from=build /usr/bin/app /app

CMD ["/app"]
