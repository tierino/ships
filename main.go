package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"example.com/decoder"
	"example.com/mapper"
	"example.com/producer"
	"example.com/tcpclient"
	"github.com/BertoldVdb/go-ais"
)

var address = "153.44.253.27:5631"

func main() {
	abort := make(chan os.Signal, 1)
	signal.Notify(abort, os.Interrupt)

	go func() {
		<-abort
		fmt.Println("\nAborting...")
		os.Exit(0)
	}()

	producer, err := producer.New(&producer.Config{BootstrapServers: "localhost:9092"})
	// todo properly handle producer errors
	if err != nil {
		panic(err)
	}
	defer producer.Disconnect()

	tcpclient := tcpclient.New()
	reader := tcpclient.ReadMessages(address)
	decoder := decoder.New()
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
			switch r := (*decoded).(type) {
			case ais.PositionReport:
				out, err := json.Marshal(mapper.TransformPositionReport(&r))
				if err != nil {
					fmt.Errorf("could not marshal PositionReport: %w", err)
				}
				producer.Send("position-reports", out)
			case ais.ShipStaticData:
				out, err := json.Marshal(mapper.TransformShipStaticData(&r))
				if err != nil {
					fmt.Errorf("could not marshal ShipStaticData: %w", err)
				}
				producer.Send("ship-static-data", out)
			default:
				fmt.Errorf("cannot handle %T messages", r)
			}
		}
	}
}
