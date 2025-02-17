package main

import (
	"fmt"
	"os"
	"os/signal"

	"example.com/decoder"
	"example.com/producer"
	"example.com/tcpclient"
	"example.com/types"
)

var address = "153.44.253.27:5631"

type Producer interface {
	Send(msg interface{}) error
}

var toTopic = map[string]string{
	string(types.PositionReportType):   "position-reports",
	string(types.StaticDataReportType): "static-data-reports",
}

func main() {
	abort := make(chan os.Signal, 1)
	signal.Notify(abort, os.Interrupt)

	go func() {
		<-abort
		fmt.Println("\nAborting...")
		os.Exit(0)
	}()

	producer := producer.New()
	tcpclient := tcpclient.New()
	reader := tcpclient.ReadMessages(address)
	defer tcpclient.Disconnect()

	for {
		raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("could not read string from buffer: %s\n", err.Error())
		}

		decoded, err := decoder.Decode(raw)
		if err != nil {
			fmt.Printf("could not decode raw payload: %s\n", err.Error())
		} else if decoded != nil {
			producer.Send(toTopic[string(decoded.Type)], *decoded)
		}
	}
}
