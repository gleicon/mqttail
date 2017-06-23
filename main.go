package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

func help() {
	fmt.Println("Usage: mqttail [-s ip:port] [-f] <topic filter>")
	fmt.Println("mqttail consumes and print messages from a mqtt topic.")
	fmt.Println("-s broker(s) separated by comma. Defaults to localhost:1883")
	fmt.Println("-q QoS level (0, 1, 2 - default QoS 0)")
	os.Exit(1)
}

func main() {

	QoSList := []byte{mqtt.QoS0, mqtt.QoS1, mqtt.QoS2}
	broker := flag.String("s", "localhost:1883", "broker")
	qos := flag.Int("q", 0, "qos")
	flag.Usage = help
	flag.Parse()

	if *qos > 2 {
		log.Println("Invalid qos, set to QoS 0")
		*qos = 0
	}

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("No topic filter given")
		os.Exit(-1)
	}

	topic_filter := args[0]

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	cli := client.New(&client.Options{
		ErrorHandler: func(err error) {
			log.Println(err)
		},
	})

	defer cli.Terminate()

	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  *broker,
		ClientID: []byte("mqttail"),
	})

	if err != nil {
		panic(err)
	}

	log.Println("Connected to ", *broker)

	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte(topic_filter),
				QoS:         QoSList[*qos],
				Handler: func(topicName, message []byte) {
					fmt.Println(string(topicName), string(message))
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	<-sigc

	err = cli.Unsubscribe(&client.UnsubscribeOptions{
		TopicFilters: [][]byte{
			[]byte(topic_filter),
		},
	})

	if err != nil {
		log.Printf("Error unsubscribing from %s: %s\n", topic_filter, err.Error())
	}

	if err := cli.Disconnect(); err != nil {
		panic(err)
	}
}
