package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"TinyGoHello/ioUtils"
	"time"
)

func init() {
	ioUtils.SetupSerialOutput(9600)
}

func main() {
	runCounter := 0

	for {
		runCounter++
		if runCounter >= 10 {
			runCounter = 0
			ioUtils.Println("Counter reset")
		}
		ioUtils.Println("A")
		time.Sleep(time.Millisecond * 500)
	}
}
