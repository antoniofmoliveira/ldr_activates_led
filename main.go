package main

import (
	"machine"
)

func main() {

	machine.InitADC()
	ldr := machine.ADC{Pin: machine.ADC0}
	ldr.Configure(machine.ADCConfig{})

	led := machine.D10
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	uart := machine.UART0
	uart.Configure(machine.UARTConfig{
		BaudRate: 9600,
		TX:       machine.UART_TX_PIN,
		RX:       machine.UART_RX_PIN,
	})

	for {
		v := ldr.Get()
		print(v)
		print(" ")
		if v < 15000 {
			led.Set(true)
		} else {
			led.Set(false)
		}
	}
}
