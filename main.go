package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"machine"
	"time"
)

func main() {
	runCounter := 0
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 9600,
	})
	for {
		runCounter++
		if runCounter >= 500 {
			println("runCounter reset triggered")
			runCounter = 0
			machine.Serial.Write([]byte("test"))
		}
		led.Low()
		time.Sleep(time.Millisecond * 50)
		machine.Serial.Write([]byte("test2"))
		led.High()
		time.Sleep(time.Millisecond * 50)
	}
}
