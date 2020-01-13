package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	v, _ := mem.VirtualMemory()

    // almost every return value is a struct
    fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

    // convert to JSON. String() is also implemented
	fmt.Println(v)
	
	opts := MQTT.NewClientOptions()
	opts.AddBroker("ws://34.87.113.63:1883")
	opts.SetClientID("MacBookPro2")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Sample Publisher Started")
	for i := 0; i < 5; i++ {
		fmt.Println("---- doing publish ----")
		token := client.Publish("something", byte(0), false, "this is the payload")	
		token.Wait()
	}

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")
}
