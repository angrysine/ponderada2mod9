package main

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var count = 0
var Time = time.Now().UnixNano() / 1000000
var msgList []string = []string{}

var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	if count == 0 {
		fmt.Printf("Qos value: %v\n", msg.Qos())
	}
	count++
	msgList = append(msgList, string(msg.Payload()[:]))


	if count == 10 {

		var current_time = time.Now().UnixNano() / 1000000
		fmt.Printf("Tempo de execução: %d ms \n", current_time-Time)
		delta := float64(current_time - Time)
		fmt.Printf("taxa de mensageria %v\n", float64(10/delta))

		fmt.Printf("Mensagem recebida: %s\n", msgList)
		defer client.Disconnect(250)
	}
}

func Subscriber() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("go_subscriber")
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	select {}
	// Bloqueia indefinidamente

}
