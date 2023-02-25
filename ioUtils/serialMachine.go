package ioUtils

import "machine"

func SetupSerialOutput(baudRate uint32) {
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: baudRate,
	})
}

func Println(toWrite string) {
	toWrite += "\r\n"
	_, _ = machine.Serial.Write([]byte(toWrite))
}
