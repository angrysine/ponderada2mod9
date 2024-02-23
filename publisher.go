package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var stop_count = 0

func Publisher() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("client.IsCOnnected() returned value: %v\n", client.IsConnected())

	for {
		if stop_count <= 10 {
			text := "Hello MQTT " + strconv.FormatInt(rand.Int64N(1000), 10)
			fmt.Printf("Publicado: %s\n", text)
			token := client.Publish("test/topic", 1, true, text)
			// fmt.Printf("Publicado: %s\n", text)
			token.Wait()
			time.Sleep(100 * time.Millisecond)
			stop_count++
		} else {
			break
		}
	}
}
