package main

import (
	"os/exec"
)

func Broker() {
	cmd := exec.Command("mosquitto", "-c", "mosquitto.conf")
	cmd.Output()

}
