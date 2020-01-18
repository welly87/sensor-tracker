package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/shirou/gopsutil/mem"
)

func main() {

	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://34.87.113.63:1883")
	opts.SetClientID("MacBookPro2")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Sample Publisher Started")

	for {
		v, _ := mem.VirtualMemory()

		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	
		token := client.Publish("something", byte(0), false, v)
		token.Wait()
		time.Sleep(1000)
	}

	client.Disconnect(250)
	fmt.Println("Sample Publisher Disconnected")
}
