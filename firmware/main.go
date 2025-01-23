package main

import (
	"fmt"
	"machine"
	"time"
)

const (
	DEBOUNCE_INTERVAL = 500 * time.Millisecond

	CMD = "OPEN"
)

func main() {
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		if button.Get() {
			fmt.Println(CMD)
			time.Sleep(DEBOUNCE_INTERVAL)
		}
	}
}
