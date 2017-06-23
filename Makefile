all: deps server

deps:
	go get -v

test:
	go test -v -bench=.

server:
	go build -v -o mqttail

clean:
	rm -f mqttail

