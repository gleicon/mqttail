deps:
	go get -v

test:
	go test -v -bench=.

all: deps
	go build -v -o mqttail

clean:
	rm -f mqttail

